package presenter

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase/response"
)

type userPresenter struct {
}

func NewUserResponseWriter() response.UserResponseWriter {
	return new(userPresenter)
}

func (p userPresenter) WriteRegisterUserResponse(user *model.User) (*response.RegisterUserResponse, error) {
	return &response.RegisterUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (p userPresenter) WriteAuthenticateUserResponse(user *model.User) (*response.AuthenticateUserResponse, error) {
	return &response.AuthenticateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
