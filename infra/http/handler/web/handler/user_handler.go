package handler

import (
	"net/http"

	"github.com/tomocy/archs/infra/http/view"
	"github.com/tomocy/archs/usecase"
)

func newUserHandler(view view.View, usecase usecase.UserUsecase) *userHandler {
	return &userHandler{
		view:    view,
		usecase: usecase,
	}
}

type userHandler struct {
	view    view.View
	usecase usecase.UserUsecase
}

func (h *userHandler) showRegistrationForm(w http.ResponseWriter, r *http.Request) {
	presenter := webPresenter(h.view, w, r)
	presenter.ShowUserRegistrationForm()
}

func (h *userHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	input, output := httpController(r), webPresenter(h.view, w, r)
	h.usecase.RegisterUser(input, output)
}
