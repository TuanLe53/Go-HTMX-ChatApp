package handlers

import (
	"net/http"
	"strings"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/TuanLe53/Go-HTMX-ChatApp/pkg/auth"
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates"
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates/components"
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

	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirm_password := c.FormValue("confirm_password")

	if !IsValidEmail(email) {
		return Render(c, components.ErrorMessage("Invalid email"))
	}

	// Check if the email exists
	_, err := models.FindUserWithEmail(email)
	if err != nil {
		return Render(c, components.ErrorMessage(err.Error()))
	}

	if strings.TrimSpace(password) != strings.TrimSpace(confirm_password) {
		return Render(c, components.ErrorMessage("Passwords don't match."))
	}

	hashedPW, err := auth.HashPw(password)
	if err != nil {
		return Render(c, components.ErrorMessage(err.Error()))
	}

	models.CreateUser(email, username, hashedPW)

	c.Response().Header().Set("hx-redirect", "/login")
	return c.NoContent(http.StatusSeeOther)
}
