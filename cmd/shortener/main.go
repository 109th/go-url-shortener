package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
)

func main() {
	err := config.ParseFlags()
	if err != nil {
		log.Fatalln(fmt.Errorf("flags parse error: %w", err))
	}

	mapStorage := types.NewMapStorage()
	s := server.NewServer(mapStorage)

	err = http.ListenAndServe(config.Addr, handlers.Router(s, config.RoutePrefix))
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(fmt.Errorf("http server error: %w", err))
	}
}
