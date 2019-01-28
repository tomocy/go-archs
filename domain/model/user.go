package model

import (
	"github.com/google/uuid"
)

type UserID string

func generateUserID() UserID {
	return UserID(uuid.New().String())
}

type User struct {
	ID       UserID
	Email    string
	Password string
}

func NewUser(email, password string) *User {
	return &User{
		ID:       generateUserID(),
		Email:    email,
		Password: password,
	}
}
