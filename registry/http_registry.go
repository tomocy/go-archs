package registry

import (
	"github.com/tomocy/archs/infra/db"
	"github.com/tomocy/archs/infra/hash"
	"github.com/tomocy/archs/infra/http/handler/web/handler"
	"github.com/tomocy/archs/infra/http/view"
	"github.com/tomocy/archs/usecase"
)

func NewHTTPRegistry() *HTTPRegistry {
	r := newWithInfras()
	r.usecase = usecase.New(r.db, r.hash)
	return r
}

func newWithInfras() *HTTPRegistry {
	return &HTTPRegistry{
		db:   db.NewMemory(),
		hash: hash.NewBcrypt(),
		view: view.NewHTML(),
	}
}

type HTTPRegistry struct {
	db      *db.Memory
	hash    *hash.Bcrypt
	view    *view.HTMLView
	usecase *usecase.Usecase
}

func (r *HTTPRegistry) NewWebHandler() *handler.Handler {
	return handler.New(r.view, r.usecase)
}
