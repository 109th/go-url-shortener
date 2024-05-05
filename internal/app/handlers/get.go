package handlers

import (
	"github.com/109th/go-url-shortener/internal/app/storage"
	"net/http"
	"strings"
)

func HandleGet(res http.ResponseWriter, req *http.Request) {
	key := strings.Trim(req.RequestURI, "/")
	value := storage.Get(key)
	if value == "" {
		http.Error(res, "400 bad request", http.StatusBadRequest)
		return
	}

	http.Redirect(res, req, value, http.StatusTemporaryRedirect)
}
