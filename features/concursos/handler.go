package concursos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/edupsousa/concursos-api/features/user"
	"github.com/edupsousa/concursos-api/platform/auth"
	"github.com/edupsousa/concursos-api/platform/httpjson"
	"github.com/gorilla/mux"
)

type Handler struct {
	concursoRepo ConcursosRepository
	userRepo     user.UserRepository
}

func NewHandler(concursoRepo ConcursosRepository, userRepo user.UserRepository) *Handler {
	return &Handler{concursoRepo: concursoRepo, userRepo: userRepo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	jwtUserRepo := &user.UserRepoJWTAdapter{UserRepository: h.userRepo}
	router.HandleFunc("/concursos", auth.WithJWTAuth(h.handleGetConcursos, jwtUserRepo)).Methods(http.MethodGet)
	router.HandleFunc("/concursos", auth.WithJWTAuth(h.handleCreateConcurso, jwtUserRepo)).Methods(http.MethodPost)
	router.HandleFunc("/concursos/{id}", auth.WithJWTAuth(h.handleGetConcurso, jwtUserRepo)).Methods(http.MethodGet)
}

func (h *Handler) handleGetConcursos(w http.ResponseWriter, r *http.Request) {
	concursos := h.concursoRepo.FindAll()

	var response []GetConcursosResponseItem
	for _, concurso := range concursos {
		response = append(response, GetConcursosResponseItem{
			ID:        concurso.ID,
			Titulo:    concurso.Titulo,
			Publicado: concurso.Publicado,
		})
	}

	httpjson.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleGetConcurso(w http.ResponseWriter, r *http.Request) {
	strID, ok := mux.Vars(r)["id"]
	if !ok {
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing concurso id"))
		return
	}

	concursoID, err := strconv.Atoi(strID)
	if err != nil {
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid concurso id"))
		return
	}

	concurso := h.concursoRepo.FindByID(concursoID)
	if concurso == nil {
		httpjson.WriteError(w, http.StatusNotFound, fmt.Errorf("concurso not found"))
		return
	}

	response := GetConcursoResponse{
		ID:        concurso.ID,
		Titulo:    concurso.Titulo,
		Publicado: concurso.Publicado,
		CreatedAt: concurso.CreatedAt,
		UpdatedAt: concurso.UpdatedAt,
	}

	httpjson.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleCreateConcurso(w http.ResponseWriter, r *http.Request) {
	var payload CreateConcursoPayload
	if err := httpjson.ParseJSON(r, &payload); err != nil {
		httpjson.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := httpjson.Validate.Struct(payload); err != nil {
		httpjson.WriteError(w, http.StatusBadRequest, err)
		return
	}

	concurso := Concurso{Titulo: payload.Titulo}
	if err := h.concursoRepo.Create(&concurso); err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpjson.WriteJSON(w, http.StatusCreated, nil)
}
