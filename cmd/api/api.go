package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/edupsousa/concursos-api/services/concursos"
	"github.com/edupsousa/concursos-api/services/user"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *sql.DB
	gdb  *gorm.DB
}

func NewAPIServer(addr string, db *sql.DB, gdb *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		gdb:  gdb,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	concursosStore := concursos.NewStore(s.gdb)
	concursosHandler := concursos.NewHandler(concursosStore, userStore)
	concursosHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
