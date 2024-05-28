package main

import (
	"errors"
	"log"
	"net/http"
	"os"

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

	//nolint:gomnd // magic allowed here ;)
	file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer func() { _ = file.Close() }()

	mapStorage, err := types.NewMapStorage(file)
	if err != nil {
		_ = file.Close()
		log.Fatalf("storage initilization error: %v", err) //nolint:gocritic // close file
	}

	s := server.NewServer(mapStorage)

	logger, err := zap.NewProduction()
	if err != nil {
		_ = file.Close()
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()

	err = http.ListenAndServe(cfg.Addr, handlers.NewRouter(s, cfg, logger))
	if !errors.Is(err, http.ErrServerClosed) {
		_ = file.Close()
		_ = logger.Sync()
		log.Fatalf("http server error: %v", err)
	}
}
