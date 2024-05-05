package handlers

import (
	"github.com/109th/go-url-shortener/internal/app/server"
	"io"
	"net/http"
)

func HandlePost(s *server.Server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
		}

		uid, err := server.GetRandomString(8)
		if err != nil {
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
		}

		s.SaveURL(uid, string(body))

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte("http://localhost:8080/" + uid))
	}
}
