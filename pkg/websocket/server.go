package websocket

import (
	"sync"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/google/uuid"
)

type WSServer struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
	Rooms      map[*ChatRoom]bool
	mu         sync.Mutex
}

func NewWsServer() *WSServer {
	return &WSServer{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
		Rooms:      make(map[*ChatRoom]bool),
	}
}

func (server *WSServer) Start() {
	for {
		select {
		case client := <-server.Register:
			server.mu.Lock()
			server.Clients[client] = true
			server.mu.Unlock()

		case client := <-server.Unregister:
			server.mu.Lock()
			delete(server.Clients, client)
			server.mu.Unlock()
			client.conn.Close()

		case message := <-server.Broadcast:
			room := server.FindRoomById(message.Target)
			if room != nil {
				room.Broadcast <- message
			}

		}
	}
}

func (server *WSServer) FindRoomById(id uuid.UUID) *ChatRoom {
	server.mu.Lock()
	defer server.mu.Unlock()

	for room := range server.Rooms {
		if room.ID == id {
			return room
		}
	}

	return nil
}

func (server *WSServer) CreateRoom(room *models.Room) *ChatRoom {
	server.mu.Lock()
	defer server.mu.Unlock()

	chatRoom := NewRoom(room)
	server.Rooms[chatRoom] = true

	go chatRoom.Start()

	return chatRoom
}
