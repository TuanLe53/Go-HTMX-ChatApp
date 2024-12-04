package models

import (
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
