package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	// ReadBufferSize:  1024,
	// WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(c echo.Context) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func ServeWS(c echo.Context, WsServer *WSServer) error {
	conn, err := Upgrade(c)
	if err != nil {
		log.Println(err)
		return err
	}

	client := newClient(conn, WsServer)

	WsServer.Register <- client
	client.Read()

	return nil
}
