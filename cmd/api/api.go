package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/edupsousa/concursos-api/services/concursos"
	"github.com/edupsousa/concursos-api/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	concursosStore := concursos.NewStore(s.db)
	concursosHandler := concursos.NewHandler(concursosStore, userStore)
	concursosHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
