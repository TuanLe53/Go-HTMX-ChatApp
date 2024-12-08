package main

import (
	"log"
	"net/http"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db"
	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/TuanLe53/Go-HTMX-ChatApp/handlers"
	custommiddleware "github.com/TuanLe53/Go-HTMX-ChatApp/middleware"
	"github.com/TuanLe53/Go-HTMX-ChatApp/pkg/websocket"
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
	db.AutoMigrate(&models.User{}, &models.Room{})

	wsServer := websocket.NewWsServer()
	go wsServer.Start()

	app.GET("/", func(c echo.Context) error {
		rooms, err := models.GetRooms(10, 0)
		if err != nil {
			log.Println("Error getting room list", err)
			c.Response().Header().Set("hx-redirect", "/")
			return c.NoContent(http.StatusSeeOther)
		}
		return handlers.Render(c, templates.Index(rooms))
	})

	authHandler := handlers.AuthHandler{}
	roomHandler := handlers.RoomHandler{}

	app.GET("/login", authHandler.LoginPage)
	app.POST("/login", authHandler.HandleLoginUser)
	app.GET("/signup", authHandler.SignUpPage)
	app.POST("/signup", authHandler.HandleSignUpUser)

	app.GET("/room/list", roomHandler.GetRoomList)

	authGroup := app.Group("")
	authGroup.Use(custommiddleware.JWTAuthMiddleware)

	authGroup.GET("/room/new", roomHandler.GetCreateRoomComponent)
	authGroup.POST("/room/new", roomHandler.CreateRoom)
	authGroup.GET("/room/:roomID", roomHandler.GetRoom)
	// authGroup.GET("/room/join/:roomID", roomHandler.JoinRoom)

	app.Logger.Fatal(app.Start(":5050"))
}
