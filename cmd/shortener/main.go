package main

import (
	"errors"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/handlers/config"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
	"net/http"
)

func main() {
	config.ParseFlags()

	mapStorage := types.NewMapStorage(50)
	s := server.NewServer(mapStorage)

	err := http.ListenAndServe(config.Addr, handlers.Router(s, config.RoutePrefix))
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
