package concursos

import (
	"database/sql"
	"fmt"

	"github.com/edupsousa/concursos-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetConcursoByID(id int) (*types.Concurso, error) {
	rows, err := s.db.Query("SELECT * FROM concursos WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	concurso := new(types.Concurso)
	for rows.Next() {
		concurso, err = scanRowIntoConcurso(rows)
		if err != nil {
			return nil, err
		}
	}

	if concurso.ID == 0 {
		return nil, fmt.Errorf("concurso not found")
	}

	return concurso, nil
}

func (s *Store) GetConcursos() ([]*types.Concurso, error) {
	rows, err := s.db.Query("SELECT * FROM concursos")
	if err != nil {
		return nil, err
	}

	concursos := make([]*types.Concurso, 0)
	for rows.Next() {
		concurso, err := scanRowIntoConcurso(rows)
		if err != nil {
			return nil, err
		}
		concursos = append(concursos, concurso)
	}

	return concursos, nil
}

func (s *Store) CreateConcurso(concurso types.Concurso) error {
	_, err := s.db.Exec("INSERT INTO concursos (titulo) VALUES (?)", concurso.Titulo)
	return err
}

func scanRowIntoConcurso(row *sql.Rows) (*types.Concurso, error) {
	var concurso types.Concurso
	if err := row.Scan(&concurso.ID, &concurso.Titulo, &concurso.CreatedAt); err != nil {
		return nil, err
	}
	return &concurso, nil
}
