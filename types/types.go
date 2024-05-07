package types

import "time"

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type User struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"emailVerified"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}