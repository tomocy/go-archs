package controller

import (
	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type UserController interface {
	RegisterUser(email, password string) (*presenter.UserPresent, error)
}

type userController struct {
	presenter presenter.UserPresenter
	usecase   usecase.UserUsecase
}

func NewUserController(presenter presenter.UserPresenter, usecase usecase.UserUsecase) UserController {
	return &userController{
		presenter: presenter,
		usecase:   usecase,
	}
}

func (c userController) RegisterUser(email, password string) (*presenter.UserPresent, error) {
	user, err := c.usecase.RegisterUser(
		request.NewRegisterUserRequest(email, password),
	)
	return c.presenter.PresentUser(user), err
}
