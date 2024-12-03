package handlers

import (
	"log"
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

func (h AuthHandler) HandleLoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if !IsValidEmail(email) {
		return Render(c, components.ErrorMessage("Invalid email"))
	}

	// Check if the email exists
	user, err := models.FindUserWithEmail(email)
	if err != nil {
		log.Println("Error finding user by email:", err)
		return Render(c, components.ErrorMessage("An error occurred, please try again later."))
	}

	// Check if user is nil (no such user)
	if user == nil {
		return Render(c, components.ErrorMessage("User with this email does not exist"))
	}

	//Compare passwords
	hashedPW := user.Password
	err = auth.CheckPw(hashedPW, password)
	if err != nil {
		return Render(c, components.ErrorMessage("Incorrect password"))
	}

	claims := auth.CreateJWTClaims(user.ID.String(), user.Name)
	token, err := auth.GenerateToken(*claims)
	if err != nil {
		log.Println("This is an error", err)
		return Render(c, components.ErrorMessage("Error logging in user."))
	}

	cookie, err := CreateCookie("access", token, 15)
	if err != nil {
		log.Println("Error creating cookie", err)
		return Render(c, components.ErrorMessage("Error logging in user."))
	}

	c.SetCookie(cookie)

	c.Response().Header().Set("hx-redirect", "/")
	return c.NoContent(http.StatusSeeOther)
}
