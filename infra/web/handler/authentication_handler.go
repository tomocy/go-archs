package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/infra/web/validator"
)

type AuthenticationHandler interface {
	AuthenticateUser(w http.ResponseWriter, r *http.Request)
}

type authenticationHandler struct {
	controller controller.AuthenticationController
}

func NewAuthenticationHandler(controller controller.AuthenticationController) AuthenticationHandler {
	return &authenticationHandler{
		controller: controller,
	}
}

func (h authenticationHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	validated, err := validator.ValidateToAuthenticateUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.controller.AuthenticateUser(w, r, validated.Email, validated.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "authenticate user: {ID: %s, Eail: %s}\n", user.ID, user.Email)
}
