package handlers

import (
	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers/middleware"
	"github.com/go-chi/chi/v5"
)

type Server interface {
	GetURL(id string) (string, error)
	SaveURL(url string) (string, error)
}

func NewRouter(s Server, cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.NewLogger())
	r.Use(middleware.NewGzipCompression())

	r.Route(cfg.RoutePrefix, func(r chi.Router) {
		r.Get("/{id}", HandleGet(s))
		r.Post("/", HandlePost(s, cfg))
		r.Post("/api/shorten", HandleShorten(s, cfg))
	})

	return r
}
