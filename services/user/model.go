package user

import (
	"time"

	"github.com/edupsousa/concursos-api/services/auth"
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

func (u *User) GetID() uint {
	return u.ID
}

type UserRepository interface {
	FindByEmail(string) *User
	FindByID(id uint) *User
	Create(*User) error
}

type UserRepoJWTAdapter struct {
	UserRepository
}

func (u *UserRepoJWTAdapter) FindByID(id uint) auth.JWTUser {
	return u.UserRepository.FindByID(id)
}
