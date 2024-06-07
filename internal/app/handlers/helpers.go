package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

func handleSaveURLError(res http.ResponseWriter, err error) {
	zap.S().Errorw(
		"save URL error: %w",
		"error", err,
	)
	http.Error(res, "500 internal server error", http.StatusInternalServerError)
}

func handleGetURLError(res http.ResponseWriter, err error) {
	zap.S().Errorw(
		"get URL error",
		"error", err,
	)
	http.Error(res, "500 internal server error", http.StatusInternalServerError)
}
