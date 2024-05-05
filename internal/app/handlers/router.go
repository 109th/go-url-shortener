package handlers

import (
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/go-chi/chi/v5"
)

func Router(s *server.Server) chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", HandleGet(s))
	r.Post("/", HandlePost(s))

	return r
}
