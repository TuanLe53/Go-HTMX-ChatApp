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

			// for client := range room.Clients {
			// 	fmt.Println(client)
			// 	notify := &Message{ClientName: client.ID, Text: "New User Joined..."}
			// 	err := client.Conn.WriteMessage(websocket.TextMessage, getTemplate("templates/notify.html", notify))
			// 	if err != nil {
			// 		fmt.Println("Error writing notify", err)
			// 		continue
			// 	}
			// }

		case client := <-room.Unregister:
			room.mu.Lock()
			delete(room.Clients, client)
			client.conn.Close()
			room.mu.Unlock()

			fmt.Println("Size of Connection Pool: ", len(room.Clients))

			// for client := range room.Clients {
			// 	notify := &Message{ClientName: client.ID, Text: "User Disconnected..."}
			// 	err := client.Conn.WriteMessage(websocket.TextMessage, getTemplate("templates/notify.html", notify))
			// 	if err != nil {
			// 		fmt.Println("Error writing notify", err)
			// 		continue
			// 	}
			// }

		case message := <-room.Broadcast:
			// room.mu.Lock()
			for client := range room.Clients {
				err := client.conn.WriteMessage(websocket.TextMessage, getTemplate("templates/components/message.html", message))
				if err != nil {
					log.Println("Error writing message", err)
					continue
				}
			}
			// room.mu.Unlock()
		}
	}
}
