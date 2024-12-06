package models

import (
	"github.com/TuanLe53/Go-HTMX-ChatApp/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"type:varchar(125);not null"`
	IsPrivate bool      `gorm:"DEFAULT:False"`
	Password  string    `gorm:"type:varchar(255)"`
}

func FindRoomByID(roomID string) *Room {

	return &Room{Name: "Hehe"}
}

func CreateChatRoom(name, password string, isPrivate bool) *Room {
	db := db.DB()

	room := Room{
		ID:        uuid.New(),
		Name:      name,
		Password:  password,
		IsPrivate: isPrivate,
	}

	db.Create(&room)

	return &room
}
