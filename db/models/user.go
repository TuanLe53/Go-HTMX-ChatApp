package models

import (
	"errors"

	"github.com/TuanLe53/Go-HTMX-ChatApp/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    string    `gorm:"type:varchar(100);not null;unique"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Password string    `gorm:"type:varchar(255);not null"`
}

func CreateUser(email, name, password string) *User {
	db := db.DB()

	user := User{Email: email, Name: name, Password: password}

	db.Create(&user)

	return &user
}

func FindUserWithEmail(email string) (*User, error) {
	db := db.DB()

	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.New("error looking for user")
	} else {
		return &user, nil
	}
}
