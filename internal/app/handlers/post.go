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

func HandlePost(s *server.Server, cfg *config.Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() { _ = req.Body.Close() }()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			handleError(res, err)
			return
		}

		uid, err := s.SaveURL(string(body))
		if err != nil {
			handleError(res, err)
			return
		}

		res.WriteHeader(http.StatusCreated)
		result, err := url.JoinPath(cfg.ServerURLPrefix, uid)
		if err != nil {
			handleError(res, err)
			return
		}
		_, _ = res.Write([]byte(result))
	}
}

func handleError(res http.ResponseWriter, err error) {
	log.Println(fmt.Errorf("save URL error: %w", err))
	http.Error(res, "500 internal server error", http.StatusInternalServerError)
}
