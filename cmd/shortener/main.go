package main

import (
	"github.com/109th/go-url-shortener/internal/app/handlers"
	srv "github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
	"net/http"
)

var Server *srv.Server

func main() {
	mux := http.NewServeMux()
	mux.Handle(`/`, http.HandlerFunc(handle))

	mapStorage := types.NewMapStorage(50)
	Server = srv.NewServer(mapStorage)

	err := http.ListenAndServe(`localhost:8080`, mux)
	if err != nil {
		panic(err)
	}
}

func handle(res http.ResponseWriter, req *http.Request) {
	var h http.HandlerFunc
	if req.Method == http.MethodPost {
		h = handlers.HandlePost(Server)
	} else {
		h = handlers.HandleGet(Server)
	}

	h(res, req)
}
