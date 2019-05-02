package server

import (
	"log"
	"net/http"

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

func (s *Server) ListenAndServe(addr string) error {
	s.registerRoutes()
	log.Printf("listen and serve on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) registerRoutes() {
	for _, r := range s.registerers {
		r.RegisterRoutes(s.router)
	}
}
