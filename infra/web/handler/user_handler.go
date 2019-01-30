package handler

import (
	"fmt"
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/infra/web/validator"
	"github.com/tomocy/archs/usecase"
)

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	controller controller.UserController
}

func NewUserHandler(controller controller.UserController) UserHandler {
	return &userHandler{
		controller: controller,
	}
}

func (h userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	validated, err := validator.ValidateToRegisterUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.controller.RegisterUser(validated.Email, validated.Password)
	if err != nil {
		switch err.(type) {
		case usecase.DuplicatedEmailError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "register user: {ID: %s, Email: %s}\n", user.ID, user.Email)
}
