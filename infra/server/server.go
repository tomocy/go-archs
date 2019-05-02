package server

import (
	"github.com/go-chi/chi"
	"github.com/tomocy/ritty-about/infra/http/route"
)

func New(rs ...route.Registerer) *Server {
	return &Server{
		router:      chi.NewRouter(),
		registerers: rs,
	}
}

type Server struct {
	router      chi.Router
	registerers []route.Registerer
}
