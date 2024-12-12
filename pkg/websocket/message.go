package websocket

type Message struct {
	Sender  *Client `json:"sender"`
	Message string  `json:"message"`
	Target  string  `json:"target"`
}
