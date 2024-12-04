package handlers

import (
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
	return Render(c, components.RoomList())
}

func (h RoomHandler) JoinRoom(c echo.Context) error {
	roomID := c.Param("roomID")

	room := models.FindRoomByID(roomID)

	return Render(c, templates.Room(room))
}
