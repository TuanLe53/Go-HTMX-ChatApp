package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func CreateCookie(name, value string, minutes int) (*http.Cookie, error) {
	if name == "" || value == "" {
		return nil, fmt.Errorf("cookie name and value cannot be empty")
	}
	if minutes <= 0 {
		return nil, fmt.Errorf("minutes must be a positive number")
	}

	return &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(time.Duration(minutes) * time.Minute),
		// Path:     path,
		// Domain:   domain,
		HttpOnly: true,
	}, nil
}
