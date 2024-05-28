package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func handleSaveURLError(res http.ResponseWriter, err error) {
	log.Println(fmt.Errorf("save URL error: %w", err))
	http.Error(res, "500 internal server error", http.StatusInternalServerError)
}

func handleGetURLError(res http.ResponseWriter, err error) {
	log.Println(fmt.Errorf("get URL error: %w", err))
	http.Error(res, "500 internal server error", http.StatusInternalServerError)
}
