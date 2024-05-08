package concursos

import (
	"time"

	"gorm.io/gorm"
)

type Concurso struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Titulo    string `gorm:"not null"`
	Publicado bool   `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type ConcursosStore interface {
	GetConcursos() []*Concurso
	GetConcursoByID(id int) *Concurso
	CreateConcurso(*Concurso) error
}

type CreateConcursoPayload struct {
	Titulo string `json:"titulo" validate:"required"`
}
