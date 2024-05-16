package handlers

import (
	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers/middleware"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(s *server.Server, cfg *config.Config, logger *zap.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.NewLogger(logger))

	r.Route(cfg.RoutePrefix, func(r chi.Router) {
		r.Get("/{id}", HandleGet(s))
		r.Post("/", HandlePost(s, cfg))
	})

	return r
}
