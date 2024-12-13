package websocket

import "github.com/google/uuid"

type Message struct {
	Sender  *Client   `json:"sender"`
	Message string    `json:"message"`
	Target  uuid.UUID `json:"target"`
}
