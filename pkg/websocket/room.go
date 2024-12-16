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
			log.Println("Client join Room which ID is", room.ID)
			log.Println("Room total members:", len(room.Clients))
			room.mu.Unlock()

			notify := &Message{Message: "User Join"}
			for client := range room.Clients {
				err := client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/notify.html", notify))
				if err != nil {
					log.Println("Error notify users")
					continue
				}
			}

		case client := <-room.Unregister:
			room.mu.Lock()
			delete(room.Clients, client)
			client.conn.Close()
			room.mu.Unlock()

			fmt.Println("Size of Connection Pool: ", len(room.Clients))

		case message := <-room.Broadcast:
			for client := range room.Clients {
				err := client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/message.html", message))
				if err != nil {
					log.Println("Error writing message", err)
					continue
				}
			}
		}
	}
}
