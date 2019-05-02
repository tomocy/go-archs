package web

import (
	"log"
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/infra/http/view"
)

func New(
	view view.View,
) *Handler {
	return &Handler{
		view: view,
	}
}

type Handler struct {
	view        view.View
	userHandler *userHandler
}

func (h *Handler) ShowUserRegistrationForm(w http.ResponseWriter, r *http.Request) {
	if err := h.view.Show(w, "user.new", nil); err != nil {
		logInternalServerError(w, "show user registration form", err)
	}
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

func logInternalServerError(w http.ResponseWriter, did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	w.WriteHeader(http.StatusInternalServerError)
}
