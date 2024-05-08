package api

import (
	"log"
	"net/http"

	"github.com/edupsousa/concursos-api/services/concursos"
	"github.com/edupsousa/concursos-api/services/user"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRepo := user.NewRepository(s.db)
	userHandler := user.NewHandler(userRepo)
	userHandler.RegisterRoutes(subrouter)

	concursosRepo := concursos.NewRepository(s.db)
	concursosHandler := concursos.NewHandler(concursosRepo, userRepo)
	concursosHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
