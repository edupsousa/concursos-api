package concursos

import (
	"log"

	concursos_model "github.com/edupsousa/concursos-api/services/concursos/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&concursos_model.Concurso{})
	return &Repository{db: db}
}

func (repo *Repository) FindByID(id int) *concursos_model.Concurso {
	var concurso concursos_model.Concurso
	err := repo.db.First(&concurso, id)
	if err.Error != nil {
		log.Printf("Error fetching concurso with id %d: %v", id, err.Error)
		return nil
	}
	return &concurso
}

func (repo *Repository) FindAll() []*concursos_model.Concurso {
	var concursos []*concursos_model.Concurso
	err := repo.db.Find(&concursos).Error
	if err != nil {
		log.Println("Error fetching concursos:", err)
		return []*concursos_model.Concurso{}
	}
	return concursos
}

func (repo *Repository) Create(concurso *concursos_model.Concurso) error {
	return repo.db.Create(&concurso).Error
}
