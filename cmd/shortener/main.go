package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("flags parse error: %v", err)
	}

	mapStorage := types.NewMapStorage()
	s := server.NewServer(mapStorage)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()

	err = http.ListenAndServe(cfg.Addr, handlers.NewRouter(s, cfg, logger))
	if !errors.Is(err, http.ErrServerClosed) {
		_ = logger.Sync()
		log.Fatalf("http server error: %v", err) //nolint:gocritic // sync logger
	}
}
