package handlers

import (
	"net/http"

	"github.com/TuanLe53/Go-HTMX-ChatApp/templates"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h AuthHandler) LoginPage(c echo.Context) error {
	return Render(c, templates.Login())
}

func (h AuthHandler) SignUpPage(c echo.Context) error {
	return Render(c, templates.SignUp())
}

func (h AuthHandler) HandleSignUpUser(c echo.Context) error {
	c.Response().Header().Set("hx-redirect", "/login")
	return c.String(http.StatusOK, "Success")
}
