package api

import (
	"log"
	"net/http"

	"github.com/edupsousa/concursos-api/features/concursos"
	"github.com/edupsousa/concursos-api/features/roles"
	"github.com/edupsousa/concursos-api/features/user"
	"github.com/edupsousa/concursos-api/platform/database"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *database.DB
}

func NewAPIServer(addr string, db *database.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	roles.NewRepository(s.db)

	userRepo := user.NewRepository(s.db)
	userHandler := user.NewHandler(userRepo)
	userHandler.RegisterRoutes(subrouter)

	concursosRepo := concursos.NewRepository(s.db)
	concursosHandler := concursos.NewHandler(concursosRepo, userRepo)
	concursosHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
