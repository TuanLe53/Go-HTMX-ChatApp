package websocket

import "github.com/google/uuid"

type Message struct {
	Sender  *Client   `json:"sender"`
	IsSelf  bool      `json:"is_self"`
	Message string    `json:"message"`
	Target  uuid.UUID `json:"target"`
	SendAt  string    `json:"send_at"`
}
