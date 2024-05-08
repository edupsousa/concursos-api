package concursos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/edupsousa/concursos-api/services/auth"
	concursos_model "github.com/edupsousa/concursos-api/services/concursos/model"
	user_model "github.com/edupsousa/concursos-api/services/user/model"
	"github.com/edupsousa/concursos-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	concursoRepo concursos_model.ConcursosRepository
	userRepo     user_model.UserRepository
}

func NewHandler(concursoRepo concursos_model.ConcursosRepository, userRepo user_model.UserRepository) *Handler {
	return &Handler{concursoRepo: concursoRepo, userRepo: userRepo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/concursos", auth.WithJWTAuth(h.handleGetConcursos, h.userRepo)).Methods(http.MethodGet)
	router.HandleFunc("/concursos", auth.WithJWTAuth(h.handleCreateConcurso, h.userRepo)).Methods(http.MethodPost)
	router.HandleFunc("/concursos/{id}", auth.WithJWTAuth(h.handleGetConcurso, h.userRepo)).Methods(http.MethodGet)
}

func (h *Handler) handleGetConcursos(w http.ResponseWriter, r *http.Request) {
	concursos := h.concursoRepo.FindAll()

	var response []concursos_model.GetConcursosResponseItem
	for _, concurso := range concursos {
		response = append(response, concursos_model.GetConcursosResponseItem{
			ID:        concurso.ID,
			Titulo:    concurso.Titulo,
			Publicado: concurso.Publicado,
		})
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleGetConcurso(w http.ResponseWriter, r *http.Request) {
	strID, ok := mux.Vars(r)["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing concurso id"))
		return
	}

	concursoID, err := strconv.Atoi(strID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid concurso id"))
		return
	}

	concurso := h.concursoRepo.FindByID(concursoID)
	if concurso == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("concurso not found"))
		return
	}

	response := concursos_model.GetConcursoResponse{
		ID:        concurso.ID,
		Titulo:    concurso.Titulo,
		Publicado: concurso.Publicado,
		CreatedAt: concurso.CreatedAt,
		UpdatedAt: concurso.UpdatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleCreateConcurso(w http.ResponseWriter, r *http.Request) {
	var payload concursos_model.CreateConcursoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	concurso := concursos_model.Concurso{Titulo: payload.Titulo}
	if err := h.concursoRepo.Create(&concurso); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
