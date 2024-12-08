package handlers

import (
	"log"
	"net/http"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates"
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates/components"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct{}

func (h RoomHandler) GetCreateRoomComponent(c echo.Context) error {
	return Render(c, components.CreateRoom())
}

func (h RoomHandler) GetRoomList(c echo.Context) error {
	rooms, err := models.GetRooms(10, 0)
	if err != nil {
		log.Println("Error getting room list", err)
		c.Response().Header().Set("hx-redirect", "/")
		return c.NoContent(http.StatusSeeOther)
	}

	return Render(c, components.RoomList(rooms))
}

func (h RoomHandler) CreateRoom(c echo.Context) error {
	roomName := c.FormValue("name")
	password := c.FormValue("password")
	private := c.FormValue("private")
	isPrivate := private == "true"

	if isPrivate && password == "" {
		return Render(c, components.ErrorMessage("Please provide password"))
	}

	room := models.CreateChatRoom(roomName, password, isPrivate)

	return Render(c, templates.Room(room))
}
