package concursos

import (
	"time"

	"gorm.io/gorm"
)

type ConcursosRepository interface {
	FindAll() []*Concurso
	FindByID(id int) *Concurso
	Create(*Concurso) error
}

type Concurso struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Titulo    string `gorm:"not null"`
	Publicado bool   `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CreateConcursoPayload struct {
	Titulo string `validate:"required"`
}

type GetConcursoResponse struct {
	ID        uint      `json:"id"`
	Titulo    string    `json:"titulo"`
	Publicado bool      `json:"publicado"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetConcursosResponseItem struct {
	ID        uint   `json:"id"`
	Titulo    string `json:"titulo"`
	Publicado bool   `json:"publicado"`
}
