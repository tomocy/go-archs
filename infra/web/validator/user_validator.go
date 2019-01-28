package validator

import (
	"errors"
	"net/http"
)

type RegisterUserRequest struct {
	Email    string
	Password string
}

func ValidateToRegisterUser(r *http.Request) (*RegisterUserRequest, error) {
	email, password := r.FormValue("email"), r.FormValue("password")
	if email == "" {
		return nil, errors.New("email should not be empty")
	}
	if password == "" {
		return nil, errors.New("password should not be empty")
	}

	return &RegisterUserRequest{
		Email:    email,
		Password: password,
	}, nil
}
