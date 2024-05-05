package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/109th/go-url-shortener/internal/app/storage"
	"io"
	"net/http"
)

func HandlePost(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "500 internal server error", http.StatusInternalServerError)
	}

	uid, err := genUID()
	if err != nil {
		http.Error(res, "500 internal server error", http.StatusInternalServerError)
	}

	storage.Set(uid, string(body))

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("http://localhost:8080/" + uid))
}

func genUID() (string, error) {
	data := make([]byte, 8)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(data), nil
}
