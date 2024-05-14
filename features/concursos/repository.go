package concursos

import (
	"log"

	"github.com/edupsousa/concursos-api/platform/database"
)

type Repository struct {
	db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	db.AutoMigrate(&Concurso{})
	return &Repository{db: db}
}

func (repo *Repository) FindByID(id int) *Concurso {
	var concurso Concurso
	err := repo.db.First(&concurso, id)
	if err.Error != nil {
		log.Printf("Error fetching concurso with id %d: %v", id, err.Error)
		return nil
	}
	return &concurso
}

func (repo *Repository) FindAll() []*Concurso {
	var concursos []*Concurso
	err := repo.db.Find(&concursos).Error
	if err != nil {
		log.Println("Error fetching concursos:", err)
		return []*Concurso{}
	}
	return concursos
}

func (repo *Repository) Create(concurso *Concurso) error {
	return repo.db.Create(&concurso).Error
}
