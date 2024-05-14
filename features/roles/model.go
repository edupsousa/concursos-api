package roles

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size:100;not null;unique"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt
}

type RoleRepository interface {
	FindAll() ([]*Role, error)
	FindByID(id uint) (*Role, error)
	Create(role *Role) error
}
