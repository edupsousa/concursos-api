package user

import (
	"log"

	"github.com/edupsousa/concursos-api/platform/database"
)

type Repository struct {
	db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	db.AutoMigrate(&User{})
	return &Repository{db: db}
}

func (repo *Repository) FindByEmail(email string) *User {
	var user User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("error getting user by email: %v", err)
		return nil
	}
	return &user
}

func (repo *Repository) FindByID(id uint) *User {
	var user User
	err := repo.db.First(&user, id).Error
	if err != nil {
		log.Printf("error getting user by id: %v", err)
		return nil
	}
	return &user
}

func (repo *Repository) Create(user *User) error {
	return repo.db.Create(user).Error
}
