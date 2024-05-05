package handlers

import (
	srv "github.com/109th/go-url-shortener/internal/app/server"
	"io"
	"net/http"
)

func HandlePost(s *srv.Server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
		}

		uid, err := srv.GetRandomString(8)
		if err != nil {
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
		}

		s.SaveURL(uid, string(body))

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte("http://localhost:8080/" + uid))
	}
}
