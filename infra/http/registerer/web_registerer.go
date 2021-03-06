package registerer

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomocy/archs/infra/http/handler/web/handler"
	"github.com/tomocy/archs/infra/http/middleware"
	"github.com/tomocy/archs/infra/http/route"
)

func NewWebRegisterer(handler *handler.Handler) *WebRegisterer {
	return &WebRegisterer{
		handler: handler,
	}
}

type WebRegisterer struct {
	handler *handler.Handler
}

func (r *WebRegisterer) RegisterRoutes(router chi.Router) {
	router.Use(middleware.RenewInvalidSession)
	router.Get("/*", http.FileServer(http.Dir("resource/public")).ServeHTTP)
	router.Get(route.Web.Route("user.new").Path, r.handler.ShowUserRegistrationForm)
	router.Post(route.Web.Route("user.create").Path, r.handler.RegisterUser)
}
