package concursos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/edupsousa/concursos-api/types"
	"github.com/edupsousa/concursos-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ConcursosStore
}

func NewHandler(store types.ConcursosStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/concursos", h.handleGetConcursos).Methods(http.MethodGet)
	router.HandleFunc("/concursos", h.handleCreateConcurso).Methods(http.MethodPost)
	router.HandleFunc("/concursos/{id}", h.handleGetConcurso).Methods(http.MethodGet)
}

func (h *Handler) handleGetConcursos(w http.ResponseWriter, r *http.Request) {
	concursos, err := h.store.GetConcursos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, concursos)
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

	concurso, err := h.store.GetConcursoByID(concursoID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, concurso)
}

func (h *Handler) handleCreateConcurso(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateConcursoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	concurso := types.Concurso{Titulo: payload.Titulo}
	if err := h.store.CreateConcurso(concurso); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
