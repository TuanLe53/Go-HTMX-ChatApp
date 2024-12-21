package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	WsServer *WSServer
	Room     *ChatRoom
	Name     string
	ID       string
}

func newClient(conn *websocket.Conn, WsServer *WSServer, room *ChatRoom, name, id string) *Client {
	return &Client{
		conn:     conn,
		WsServer: WsServer,
		Room:     room,
		Name:     name,
		ID:       id,
	}
}

func (c *Client) Read() {
	defer func() {
		c.WsServer.Unregister <- c
		c.Room.Unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Read error", err)
			break
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Fatal("Error Unmarshalling message")
		}

		message.Target = c.Room.ID
		message.Sender = c

		currentTime := time.Now()
		message.SendAt = currentTime.Format("15:04 02/01/2006")

		c.WsServer.Broadcast <- &message
	}
}
