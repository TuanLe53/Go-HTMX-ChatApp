package handlers

import (
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h AuthHandler) LoginPage(c echo.Context) error {
	return Render(c, templates.Login())
}
