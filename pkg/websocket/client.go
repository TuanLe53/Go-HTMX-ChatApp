package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	WsServer *WSServer
	Room     *ChatRoom
}

func newClient(conn *websocket.Conn, WsServer *WSServer, room *ChatRoom) *Client {
	return &Client{
		conn:     conn,
		WsServer: WsServer,
		Room:     room,
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

		c.WsServer.Broadcast <- &message
	}
}
