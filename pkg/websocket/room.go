package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	ID         uuid.UUID
	Name       string
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan *Message
	mu         sync.Mutex
}

func NewRoom(room *models.Room) *ChatRoom {
	return &ChatRoom{
		ID:         room.ID,
		Name:       room.Name,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *Message),
	}
}

func (room *ChatRoom) Start() {
	for {
		select {
		case client := <-room.Register:
			room.mu.Lock()
			room.Clients[client] = true
			room.mu.Unlock()

			notify := &Message{Message: fmt.Sprintf("%s has joined room.", client.Name)}
			members := getRoomMembers(room)

			for client := range room.Clients {
				err := client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/notify.html", notify))
				if err != nil {
					log.Println("Error notify users")
					continue
				}
				err = client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/member.html", members))
				if err != nil {
					log.Println("Error add user to member list")
					continue
				}
			}

		case client := <-room.Unregister:
			room.mu.Lock()
			delete(room.Clients, client)
			client.conn.Close()
			room.mu.Unlock()

			notify := &Message{Message: fmt.Sprintf("%s left room.", client.Name)}
			members := getRoomMembers(room)

			if len(members.Clients) == 0 {
				err := models.DeleteRoom(room.ID)
				if err != nil {
					log.Printf("Error deleting room from database: %v", err)
				}
				break
			}

			for client := range room.Clients {
				err := client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/notify.html", notify))
				if err != nil {
					log.Println("Error notify users")
					continue
				}
				err = client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/member.html", members))
				if err != nil {
					log.Println("Error remove user to member list")
					continue
				}
			}

		case message := <-room.Broadcast:
			for client := range room.Clients {
				message.IsSelf = client == message.Sender
				err := client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/message.html", message))
				if err != nil {
					log.Println("Error writing message", err)
					continue
				}
			}
		}
	}
}

func getRoomMembers(room *ChatRoom) struct{ Clients []*Client } {
	clients := []*Client{}
	for client := range room.Clients {
		clients = append(clients, client)
	}

	members := struct {
		Clients []*Client
	}{
		Clients: clients,
	}

	return members
}
