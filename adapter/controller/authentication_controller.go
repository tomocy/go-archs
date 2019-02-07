package controller

import (
	"net/http"

	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type AuthenticationController interface {
	AuthenticateUser(w http.ResponseWriter, r *http.Request, email, password string) (*presenter.AuthenticUserPresent, error)
	GetAuthenticUserID(r *http.Request) string
}

type authenticationController struct {
	presenter presenter.AuthenticationPresenter
	usecase   usecase.AuthenticationUsecase
}

func NewAuthenticationController(presenter presenter.AuthenticationPresenter, usecase usecase.AuthenticationUsecase) AuthenticationController {
	return &authenticationController{
		presenter: presenter,
		usecase:   usecase,
	}
}

func (c authenticationController) AuthenticateUser(w http.ResponseWriter, r *http.Request, email, password string) (*presenter.AuthenticUserPresent, error) {
	user, err := c.usecase.AuthenticateUser(
		request.NewAuthenticateUserRequest(w, r, email, password),
	)
	return c.presenter.PresentAuthenticUser(user), err
}

func (c authenticationController) GetAuthenticUserID(r *http.Request) string {
	return string(c.usecase.GetAuthenticUserID(
		request.NewGetAuthenticUserIDRequest(r),
	))
}
