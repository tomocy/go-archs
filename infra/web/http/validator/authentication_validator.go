package validator

import "net/http"

type AuthenticateUserRequest struct {
	Email    string
	Password string
}

func ValidateToAuthenticateUser(r *http.Request) (*AuthenticateUserRequest, error) {
	email, password := r.FormValue("email"), r.FormValue("password")
	if email == "" {
		return nil, newEmptyError("email")
	}
	if password == "" {
		return nil, newEmptyError("password")
	}

	return &AuthenticateUserRequest{
		Email:    email,
		Password: password,
	}, nil
}
