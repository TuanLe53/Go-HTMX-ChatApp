package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
		log.Println("Error from Upgrade", err)
		return nil, err
	}

	return conn, nil
}

func ServeWS(c echo.Context, WsServer *WSServer) error {
	roomID, err := uuid.Parse(c.QueryParam("roomID"))
	if err != nil {
		log.Println("Invalid room id")
		return err
	}

	room := WsServer.FindRoomById(roomID)
	if room == nil {
		return err
	}

	conn, err := Upgrade(c)
	if err != nil {
		log.Println("Error from ServeWs", err)
		return err
	}

	userClaims := c.Get("user")
	user, ok := userClaims.(jwt.MapClaims)
	if !ok {
		log.Println("Failed to retrieve user claims")
		return fmt.Errorf("invalid user claims")
	}

	username, ok := user["name"].(string)
	if !ok {
		log.Println("Invalid username type in claims")
		return fmt.Errorf("invalid username type")
	}

	userID, ok := user["id"].(string)
	if !ok {
		log.Println("Invalid user ID type in claims")
		return fmt.Errorf("invalid user ID type")
	}

	client := newClient(conn, WsServer, room, username, userID)

	WsServer.Register <- client
	room.Register <- client

	client.Read()

	return nil
}
