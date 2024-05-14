package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/go-chi/chi/v5"
)

func HandleGet(s *server.Server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		key := chi.URLParam(req, "id")
		value, err := s.GetURL(key)
		if err != nil {
			if errors.Is(err, server.ErrNotFound) {
				http.Error(res, "400 bad request", http.StatusBadRequest)
				return
			}

			log.Println(fmt.Errorf("get URL error: %w", err))
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, value, http.StatusTemporaryRedirect)
	}
}
