package middleware

import (
	"log"
	"net/http"

	"github.com/TuanLe53/Go-HTMX-ChatApp/handlers"
	"github.com/TuanLe53/Go-HTMX-ChatApp/pkg/auth"
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates/components"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access")
		if err != nil {
			if err == http.ErrNoCookie {
				// Cookie not found
				log.Println("No access cookie found")
				return handlers.Render(c, components.AccessDenied())
			} else {
				// other errors
				log.Println("Error retrieving cookie:", err)
				return handlers.Render(c, components.ErrorMessage("Something went wrong. Please try again later"))
			}
		}

		token, err := auth.ValidateToken(cookie.Value)
		if err != nil {
			log.Println("Invalid token: ", err)
			c.Response().Header().Set("hx-redirect", "/login")
			return c.NoContent(http.StatusSeeOther)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Invalid token claims: ")
			c.Response().Header().Set("hx-redirect", "/login")
			return c.NoContent(http.StatusSeeOther)
		}
		c.Set("user", claims)

		return next(c)
	}
}
