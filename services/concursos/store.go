package concursos

import (
	"log"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	db.AutoMigrate(&Concurso{})
	return &Store{db: db}
}

func (s *Store) GetConcursoByID(id int) *Concurso {
	var concurso Concurso
	err := s.db.First(&concurso, id)
	if err.Error != nil {
		log.Printf("Error fetching concurso with id %d: %v", id, err.Error)
		return nil
	}
	return &concurso
}

func (s *Store) GetConcursos() []*Concurso {
	var concursos []*Concurso
	err := s.db.Find(&concursos).Error
	if err != nil {
		log.Println("Error fetching concursos:", err)
		return []*Concurso{}
	}
	return concursos
}

func (s *Store) CreateConcurso(concurso *Concurso) error {
	return s.db.Create(&concurso).Error
}
