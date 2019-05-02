package registerer

import (
	"github.com/go-chi/chi"
	"github.com/tomocy/archs/infra/http/handler/web"
)

func NewWebRegisterer(handler *web.Handler) *WebRegisterer {
	return &WebRegisterer{
		handler: handler,
	}
}

type WebRegisterer struct {
	handler *web.Handler
}

func (r *WebRegisterer) RegisterRoutes(router chi.Router) {
	router.Post("/users", r.handler.RegisterUser)
}
