package request

import "net/http"

type RegisterUserRequest struct {
	Email    string
	Password string
}

func NewRegisterUserRequest(email, password string) *RegisterUserRequest {
	return &RegisterUserRequest{
		Email:    email,
		Password: password,
	}
}

type AuthenticateUserRequest struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Email          string
	Password       string
}

func NewAuthenticateUserRequest(w http.ResponseWriter, r *http.Request, email, password string) *AuthenticateUserRequest {
	return &AuthenticateUserRequest{
		ResponseWriter: w,
		Request:        r,
		Email:          email,
		Password:       password,
	}
}
