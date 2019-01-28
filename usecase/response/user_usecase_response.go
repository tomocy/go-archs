package response

import "github.com/tomocy/archs/domain/model"

type UserResponseWriter interface {
	WriteRegisterUserResponse(user *model.User) (*RegisterUserResponse, error)
}

type RegisterUserResponse struct {
	ID    model.UserID
	Email string
}
