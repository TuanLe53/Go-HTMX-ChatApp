package main

import (
	"log"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db"
	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/TuanLe53/Go-HTMX-ChatApp/handlers"
	custommiddleware "github.com/TuanLe53/Go-HTMX-ChatApp/middleware"
	"github.com/TuanLe53/Go-HTMX-ChatApp/templates"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Static("/static", "assets")

	app.Use(echomiddleware.Logger())
	app.Use(echomiddleware.Recover())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	db := db.DB()
	db.AutoMigrate(&models.User{})

	app.GET("/", func(c echo.Context) error {
		return handlers.Render(c, templates.Index())
	})

	authHandler := handlers.AuthHandler{}
	app.GET("/login", authHandler.LoginPage)
	app.POST("/login", authHandler.HandleLoginUser)
	app.GET("/signup", authHandler.SignUpPage)
	app.POST("/signup", authHandler.HandleSignUpUser)

	authGroup := app.Group("")
	authGroup.Use(custommiddleware.JWTAuthMiddleware)

	roomHandler := handlers.RoomHandler{}
	authGroup.GET("/room/new", roomHandler.GetCreateRoomComponent)
	authGroup.GET("/room/list", roomHandler.GetRoomList)

	app.Logger.Fatal(app.Start(":5050"))
}
