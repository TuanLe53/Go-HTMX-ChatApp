package models

import (
	"errors"

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

func FindRoomByID(roomID string) (*Room, error) {
	db := db.DB()

	var room Room
	result := db.Where("id = ?", roomID).First(&room)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.New("error looking for room")
	}

	return &room, nil
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