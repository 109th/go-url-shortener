package main

import (
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
	"net/http"
)

func main() {
	mapStorage := types.NewMapStorage(50)
	s := server.NewServer(mapStorage)

	err := http.ListenAndServe(`localhost:8080`, handlers.Router(s))
	if err != nil {
		panic(err)
	}
}
