package controller

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type UserController interface {
	RegisterUser(email, password string) (*model.User, error)
}

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return &userController{
		usecase: usecase,
	}
}

func (c userController) RegisterUser(email, password string) (*model.User, error) {
	return c.usecase.RegisterUser(
		request.NewRegisterUserRequest(email, password),
	)
}
