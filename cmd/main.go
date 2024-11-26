package main

import (
	"log"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db"
	"github.com/TuanLe53/Go-HTMX-ChatApp/db/models"
	"github.com/TuanLe53/Go-HTMX-ChatApp/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	db := db.DB()
	db.AutoMigrate(&models.User{})

	authHandler := handlers.AuthHandler{}
	app.GET("/login", authHandler.LoginPage)
	app.GET("/signup", authHandler.SignUpPage)
	app.POST("/signup", authHandler.HandleSignUpUser)

	app.Logger.Fatal(app.Start(":5050"))
}
