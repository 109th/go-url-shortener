package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/server"
)

func HandlePost(s *server.Server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() { _ = req.Body.Close() }()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
			return
		}

		uid, err := s.SaveURL(string(body))
		if err != nil {
			log.Println(fmt.Errorf("save URL error: %w", err))
			http.Error(res, "500 internal server error", http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusCreated)
		result, _ := url.JoinPath(config.ServerURLPrefix, uid)
		_, _ = res.Write([]byte(result))
	}
}
