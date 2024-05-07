package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	token, err := CreateJWT(1)
	if err != nil {
		t.Errorf("error creating JWT: %s", err)
	}
	if token == "" {
		t.Error("expected token to be not empty")
	}
}
