package handlers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/109th/go-url-shortener/internal/app/config"
)

func HandlePost(s Server, cfg *config.Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() { _ = req.Body.Close() }()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			handleSaveURLError(res, err)
			return
		}

		uid, err := s.SaveURL(string(body))
		if err != nil {
			handleSaveURLError(res, err)
			return
		}

		res.WriteHeader(http.StatusCreated)
		result, err := url.JoinPath(cfg.ServerURLPrefix, uid)
		if err != nil {
			handleSaveURLError(res, err)
			return
		}
		_, _ = res.Write([]byte(result))
	}
}
