package web

import (
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/usecase"
)

func New(
	userUsecase usecase.UserUsecase,
) *Handler {
	return &Handler{
		userHandler: newUserHandler(userUsecase),
	}
}

type Handler struct {
	userHandler *userHandler
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	h.userHandler.registerUser(w, r)
}

func httpController(r *http.Request) *controller.HTTPController {
	return controller.NewHTTPController(r)
}

func httpPresenter(w http.ResponseWriter, r *http.Request) *presenter.HTTPPresenter {
	return presenter.NewHTTPPresenter(w, r)
}
