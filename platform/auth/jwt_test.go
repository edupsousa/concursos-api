package auth

import "testing"

type mockUser struct {
	ID     uint
	RoleID uint
}

func (u *mockUser) GetID() uint {
	return u.ID
}

func (u *mockUser) GetRoleID() uint {
	return u.RoleID
}

func TestCreateJWT(t *testing.T) {
	u := &mockUser{ID: 1, RoleID: 1}
	token, err := CreateJWT(u)
	if err != nil {
		t.Errorf("error creating JWT: %s", err)
	}
	if token == "" {
		t.Error("expected token to be not empty")
	}
}
