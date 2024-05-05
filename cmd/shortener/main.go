package main

import (
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/storage"
	"net/http"
)

func init() {
	storage.Init(50)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(`/`, http.HandlerFunc(handle))

	err := http.ListenAndServe(`localhost:8080`, mux)
	if err != nil {
		panic(err)
	}
}

func handle(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		handlers.HandlePost(res, req)
	} else {
		handlers.HandleGet(res, req)
	}
}
