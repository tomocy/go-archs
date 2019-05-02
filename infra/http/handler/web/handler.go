package web

import (
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/adapter/presenter"
)

func httpController(r *http.Request) *controller.HTTPController {
	return controller.NewHTTPController(r)
}

func httpPresenter(w http.ResponseWriter, r *http.Request) *presenter.HTTPPresenter {
	return presenter.NewHTTPPresenter(w, r)
}
