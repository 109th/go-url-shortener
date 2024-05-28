package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

var compatibleTypes = map[string]bool{
	"application/json": true,
	"text/html":        true,
}

type gzipResponseWriter struct {
	http.ResponseWriter
	compressible    bool
	isHeaderWritten bool
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	var writer io.Writer

	if w.compressible {
		cw := gzip.NewWriter(w.ResponseWriter)
		defer func() { _ = cw.Close() }()
		writer = cw
	} else {
		writer = w.ResponseWriter
	}

	return writer.Write(b) //nolint:wrapcheck // decorator
}

func (w *gzipResponseWriter) WriteHeader(statusCode int) {
	if !w.isHeaderWritten {
		w.compressible = w.isCompatible()

		if w.compressible {
			w.Header().Set("Content-Encoding", "gzip")
		}
	}

	w.isHeaderWritten = true
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *gzipResponseWriter) isCompatible() bool {
	contentType := w.Header().Get("Content-Type")
	if idx := strings.Index(contentType, ";"); idx >= 0 {
		contentType = contentType[0:idx]
	}

	return compatibleTypes[contentType]
}

func NewGzipCompression(logger *zap.Logger) func(http.Handler) http.Handler {
	slog := logger.Sugar()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ow := w

			contentEncoding := r.Header.Get("Content-Encoding")
			if strings.Contains(contentEncoding, "gzip") {
				reader, err := gzip.NewReader(r.Body)
				if err != nil {
					slog.Errorw("failed to decompress request", "error", err)
					http.Error(w, "gzip middleware error", http.StatusInternalServerError)

					return
				}

				r.Body = reader
			}

			acceptEncoding := r.Header.Get("Accept-Encoding")
			if strings.Contains(acceptEncoding, "gzip") {
				cw := &gzipResponseWriter{
					ResponseWriter: w,
				}
				ow = cw
			}

			next.ServeHTTP(ow, r)
		})
	}
}
