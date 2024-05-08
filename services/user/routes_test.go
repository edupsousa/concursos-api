package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	user_model "github.com/edupsousa/concursos-api/services/user/model"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userRepo := &mockUserRepo{}
	handler := NewHandler(userRepo)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := user_model.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "abc",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register a new user", func(t *testing.T) {
		payload := user_model.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john_doe@acme.com",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserRepo struct{}

func (m *mockUserRepo) FindByEmail(email string) *user_model.User {
	return nil
}

func (m *mockUserRepo) FindByID(id int) *user_model.User {
	return nil
}

func (m *mockUserRepo) Create(user *user_model.User) error {
	return nil
}
