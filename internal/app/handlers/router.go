package handlers

import (
	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/go-chi/chi/v5"
)

func Router(s *server.Server, cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Route(cfg.RoutePrefix, func(r chi.Router) {
		r.Get("/{id}", HandleGet(s))
		r.Post("/", HandlePost(s, cfg))
	})

	return r
}
