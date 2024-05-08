package user_model

import (
	"time"

	"gorm.io/gorm"
)

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	Email         string `gorm:"not null;unique"`
	EmailVerified bool   `gorm:"not null;default:false"`
	Password      string `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type UserStore interface {
	GetUserByEmail(string) *User
	GetUserByID(id int) *User
	CreateUser(*User) error
}
