package presenter

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase/response"
)

func NewUserUsecaseResponser() response.UserUsecaseResponser {
	return new(userPresenter)
}

type userPresenter struct {
}

func (p userPresenter) ResponseUser(user *model.User) *response.UserResponse {
	return response.NewUserResponse(string(user.ID), user.Email)
}
