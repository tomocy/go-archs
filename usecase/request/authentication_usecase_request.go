package request

import "net/http"

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
