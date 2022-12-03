package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/internal/domain/author"
)

type Server struct {
	router *chi.Mux
	client *ent.Client
}

func NewServer(client *ent.Client) *Server {
	return &Server{
		router: chi.NewRouter(),
		client: client,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) RouteRegister() {
	// Middleware
	s.router.Use(middleware.Logger)

	// Author
	authorServer := author.NewServer(s.client)
	authorHandler := author.NewHandler(authorServer)
	s.router.Mount("/authors", author.Routes(authorHandler))
}
