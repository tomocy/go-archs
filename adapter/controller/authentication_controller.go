package controller

import (
	"net/http"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type AuthenticationController interface {
	AuthenticateUser(w http.ResponseWriter, r *http.Request, email, password string) (*model.User, error)
}

type authenticationController struct {
	usecase usecase.AuthenticationUsecase
}

func NewAuthenticationController(usecase usecase.AuthenticationUsecase) AuthenticationController {
	return &authenticationController{
		usecase: usecase,
	}
}

func (c authenticationController) AuthenticateUser(w http.ResponseWriter, r *http.Request, email, password string) (*model.User, error) {
	return c.usecase.AuthenticateUser(
		request.NewAuthenticateUserRequest(w, r, email, password),
	)
}
