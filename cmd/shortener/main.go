package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("flags parse error: %v", err)
	}

	mapStorage := types.NewMapStorage()
	s := server.NewServer(mapStorage)

	err = http.ListenAndServe(cfg.Addr, handlers.Router(s, cfg))
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("http server error: %v", err)
	}
}
