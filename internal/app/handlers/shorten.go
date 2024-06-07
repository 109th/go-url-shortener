package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/109th/go-url-shortener/internal/app/config"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

func HandleShorten(s Server, cfg *config.Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var body ShortenRequest
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			http.Error(res, "400 bad request", http.StatusBadRequest)
		}

		uid, err := s.SaveURL(body.URL)
		if err != nil {
			handleSaveURLError(res, err)
			return
		}

		result, err := url.JoinPath(cfg.ServerURLPrefix, uid)
		if err != nil {
			handleSaveURLError(res, err)
			return
		}

		response := &ShortenResponse{
			Result: result,
		}

		res.Header().Add("content-type", "application/json")
		res.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(res).Encode(response)
		if err != nil {
			handleSaveURLError(res, err)
			return
		}
	}
}
