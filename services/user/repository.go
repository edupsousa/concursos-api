package user

import (
	"log"

	user_model "github.com/edupsousa/concursos-api/services/user/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&user_model.User{})
	return &Repository{db: db}
}

func (repo *Repository) FindByEmail(email string) *user_model.User {
	var user user_model.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("error getting user by email: %v", err)
		return nil
	}
	return &user
}

func (repo *Repository) FindByID(id int) *user_model.User {
	var user user_model.User
	err := repo.db.First(&user, id).Error
	if err != nil {
		log.Printf("error getting user by id: %v", err)
		return nil
	}
	return &user
}

func (repo *Repository) Create(user *user_model.User) error {
	return repo.db.Create(user).Error
}
