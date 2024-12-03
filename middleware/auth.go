package middleware

import (
	"log"
	"net/http"

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
		log.Printf("Cookie: Name=%s, Value=%s", cookie.Name, cookie.Value)

		return next(c)
	}
}
