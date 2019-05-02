package registerer

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomocy/archs/infra/http/handler/web"
	"github.com/tomocy/archs/infra/http/route"
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
	router.Get("/*", http.FileServer(http.Dir("resource/public")).ServeHTTP)
	router.Get(route.Web.Route("user.new").Path, r.handler.ShowUserRegistrationForm)
	router.Post(route.Web.Route("user.create").Path, r.handler.RegisterUser)
}
