package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("flags parse error: %v", err)
	}

	s, err := storage.NewStorage(cfg)
	if err != nil {
		log.Fatalf("storage initilization error: %v", err)
	}
	closeStorage := func() {
		err := s.Close()
		if err != nil {
			log.Fatalf("storage close error: %v", err)
		}
	}
	defer closeStorage()

	srv := server.NewServer(s)

	logger, err := zap.NewProduction()
	if err != nil {
		closeStorage()
		log.Fatalf("can't initialize zap logger: %v", err) //nolint:gocritic // close storage
	}
	zap.ReplaceGlobals(logger)
	defer func() { _ = logger.Sync() }()

	err = http.ListenAndServe(cfg.Addr, handlers.NewRouter(srv, cfg))
	if !errors.Is(err, http.ErrServerClosed) {
		closeStorage()
		_ = logger.Sync()
		log.Fatalf("http server error: %v", err)
	}
}
