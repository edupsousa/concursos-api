package user

import (
	"log"

	user_model "github.com/edupsousa/concursos-api/services/user/model"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	db.AutoMigrate(&user_model.User{})
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) *user_model.User {
	var user user_model.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("error getting user by email: %v", err)
		return nil
	}
	return &user
}

func (s *Store) GetUserByID(id int) *user_model.User {
	var user user_model.User
	err := s.db.First(&user, id).Error
	if err != nil {
		log.Printf("error getting user by id: %v", err)
		return nil
	}
	return &user
}

func (s *Store) CreateUser(user *user_model.User) error {
	return s.db.Create(user).Error
}
