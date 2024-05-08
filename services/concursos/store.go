package concursos

import (
	"log"

	concursos_model "github.com/edupsousa/concursos-api/services/concursos/model"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	db.AutoMigrate(&concursos_model.Concurso{})
	return &Store{db: db}
}

func (s *Store) GetConcursoByID(id int) *concursos_model.Concurso {
	var concurso concursos_model.Concurso
	err := s.db.First(&concurso, id)
	if err.Error != nil {
		log.Printf("Error fetching concurso with id %d: %v", id, err.Error)
		return nil
	}
	return &concurso
}

func (s *Store) GetConcursos() []*concursos_model.Concurso {
	var concursos []*concursos_model.Concurso
	err := s.db.Find(&concursos).Error
	if err != nil {
		log.Println("Error fetching concursos:", err)
		return []*concursos_model.Concurso{}
	}
	return concursos
}

func (s *Store) CreateConcurso(concurso *concursos_model.Concurso) error {
	return s.db.Create(&concurso).Error
}
