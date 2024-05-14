package user

import (
	"fmt"
	"net/http"

	"github.com/edupsousa/concursos-api/features/auth"
	"github.com/edupsousa/concursos-api/platform/httpjson"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	userRepo UserRepository
}

func NewHandler(userRepo UserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload
	if err := httpjson.ParseJSON(r, &payload); err != nil {
		httpjson.WriteError(w, http.StatusBadRequest, err)
	}

	if err := httpjson.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	u := h.userRepo.FindByEmail(payload.Email)
	if u == nil {
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	if !auth.ComparePassword(u.Password, payload.Password) {
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	token, err := auth.CreateJWT(u)
	if err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpjson.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := httpjson.ParseJSON(r, &payload); err != nil {
		httpjson.WriteError(w, http.StatusBadRequest, err)
	}

	if err := httpjson.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	// TODO: Replace with count user by email
	user := h.userRepo.FindByEmail(payload.Email)
	if user != nil {
		httpjson.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.userRepo.Create(&User{
		FirstName:     payload.FirstName,
		LastName:      payload.LastName,
		Email:         payload.Email,
		EmailVerified: false,
		Password:      hashedPassword,
	})
	if err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpjson.WriteJSON(w, http.StatusCreated, nil)
}
