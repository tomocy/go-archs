package controller

import (
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
	"github.com/tomocy/archs/usecase/response"
)

type UserController interface {
	RegisterUser(email, password string) (*response.UserResponse, error)
}

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return &userController{
		usecase: usecase,
	}
}

func (c userController) RegisterUser(email, password string) (*response.UserResponse, error) {
	return c.usecase.RegisterUser(
		request.NewRegisterUserRequest(email, password),
	)
}
