package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("DSN")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
}

func DB() *gorm.DB {
	return db
}
