package controller

import (
	"net/http"

	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
	"github.com/tomocy/archs/usecase/response"
)

type AuthenticationController interface {
	AuthenticateUser(w http.ResponseWriter, r *http.Request, email, password string) (*response.UserResponse, error)
	GetAuthenticUserID(r *http.Request) string
}

type authenticationController struct {
	usecase usecase.AuthenticationUsecase
}

func NewAuthenticationController(usecase usecase.AuthenticationUsecase) AuthenticationController {
	return &authenticationController{
		usecase: usecase,
	}
}

func (c authenticationController) AuthenticateUser(w http.ResponseWriter, r *http.Request, email, password string) (*response.UserResponse, error) {
	return c.usecase.AuthenticateUser(
		request.NewAuthenticateUserRequest(w, r, email, password),
	)
}

func (c authenticationController) GetAuthenticUserID(r *http.Request) string {
	return string(c.usecase.GetAuthenticUserID(
		request.NewGetAuthenticUserIDRequest(r),
	))
}
