package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "a@a.ru",
		Password: "password",
	}
}
