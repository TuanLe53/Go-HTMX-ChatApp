package handlers

import (
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates/components"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct{}

func (h RoomHandler) GetCreateRoomComponent(c echo.Context) error {
	return Render(c, components.CreateRoom())
}
