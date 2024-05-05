package handlers

import (
	"errors"
	srv "github.com/109th/go-url-shortener/internal/app/server"
	"net/http"
	"strings"
)

func HandleGet(s *srv.Server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		key := strings.Trim(req.RequestURI, "/")
		value, err := s.GetURL(key)
		if err != nil {
			if errors.Is(err, srv.ErrNotFound) {
				http.Error(res, "400 bad request", http.StatusBadRequest)
				return
			}

			http.Error(res, "500 internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, value, http.StatusTemporaryRedirect)
	}
}
