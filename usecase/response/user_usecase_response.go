package response

import "github.com/tomocy/archs/domain/model"

type UserResponseWriter interface {
	WriteRegisterUserResponse(user *model.User) (*RegisterUserResponse, error)
	WriteAuthenticateUserResponse(user *model.User) (*AuthenticateUserResponse, error)
}

type RegisterUserResponse struct {
	ID    model.UserID
	Email string
}

type AuthenticateUserResponse struct {
	ID    model.UserID
	Email string
}
