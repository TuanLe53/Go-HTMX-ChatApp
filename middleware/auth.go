package middleware

import (
	"log"
	"net/http"

	"github.com/TuanLe53/Go-HTMX-ChatApp/pkg/auth"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("MIDDLEWARE CALL")
		cookie, err := c.Cookie("access")
		if err != nil {
			if err == http.ErrNoCookie {
				// Cookie not found
				log.Println("No access cookie found")
				c.Response().Header().Set("hx-redirect", "/login")
				return c.NoContent(http.StatusSeeOther)
			} else {
				// Some other error
				log.Println("Error retrieving cookie:", err)
				c.Response().Header().Set("hx-redirect", "/login")
				return c.NoContent(http.StatusSeeOther)
			}
		}

		// Log the cookie value
		log.Printf("Token available")
		_, err = auth.ValidateToken(cookie.Value)
		if err != nil {
			log.Println("Invalid token: ", err)
			c.Response().Header().Set("hx-redirect", "/login")
			return c.NoContent(http.StatusSeeOther)
		}

		return next(c)
	}
}
