package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size

	return size, err //nolint:wrapcheck // decorator
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func NewLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	sugar := logger.Sugar()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			uri := r.RequestURI
			method := r.Method

			responseData := &responseData{}
			lw := loggingResponseWriter{
				ResponseWriter: w,
				responseData:   responseData,
			}
			next.ServeHTTP(&lw, r)

			duration := time.Since(start)

			sugar.Infoln(
				"method", method,
				"uri", uri,
				"duration", duration,
				"status", responseData.status,
				"size", responseData.size,
			)
		})
	}
}
