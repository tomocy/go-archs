package web

import (
	"net/http"

	"github.com/tomocy/archs/usecase"
)

func newUserHandler(usecase usecase.UserUsecase) *userHandler {
	return &userHandler{
		usecase: usecase,
	}
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func (h *userHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	input, output := httpController(r), webPresenter(w, r)
	h.usecase.RegisterUser(input, output)
}
