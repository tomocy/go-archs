package registry

import (
	"github.com/tomocy/archs/infra/http/handler/web"
	"github.com/tomocy/archs/infra/http/view"
)

func NewHTTPRegistry() *HTTPRegistry {
	return &HTTPRegistry{
		view: view.NewHTML(),
	}
}

type HTTPRegistry struct {
	view *view.HTMLView
}

func (r *HTTPRegistry) NewWebHandler() *web.Handler {
	return web.New(r.view)
}
