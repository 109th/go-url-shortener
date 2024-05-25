package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/handlers/internal/mockery"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHandleShorten(t *testing.T) {
	type want struct {
		statusCode int
		response   *handlers.ShortenResponse
	}
	type configs struct {
		baseURLPrefix string
		routePrefix   string
	}
	tests := []struct {
		name    string
		request *handlers.ShortenRequest
		cfg     configs
		want    want
	}{
		{
			name: "test create redirect url",
			request: &handlers.ShortenRequest{
				URL: "https://example.com",
			},
			want: want{
				statusCode: 201,
				response: &handlers.ShortenResponse{
					Result: "http://localhost:8080/mock-generated-id",
				},
			},
		},
		{
			name: "test create redirect url with prefix",
			request: &handlers.ShortenRequest{
				URL: "https://example.com",
			},
			cfg: configs{
				baseURLPrefix: "http://localhost:8081/prefix",
				routePrefix:   "/prefix",
			},
			want: want{
				statusCode: 201,
				response: &handlers.ShortenResponse{
					Result: "http://localhost:8081/prefix/mock-generated-id",
				},
			},
		},
	}

	logger, _ := zap.NewDevelopment()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// config setup
			cfg := &config.Config{
				RoutePrefix:     "/",
				ServerURLPrefix: "http://localhost:8080",
			}
			if tt.cfg.routePrefix != "" {
				cfg.RoutePrefix = tt.cfg.routePrefix
			}

			if tt.cfg.baseURLPrefix != "" {
				cfg.ServerURLPrefix = tt.cfg.baseURLPrefix
			}

			// restore default config
			defer func() {
				cfg.RoutePrefix = "/"
				cfg.ServerURLPrefix = "http://localhost:8080"
			}()

			srv := mockery.NewMockServer(t)
			srv.EXPECT().SaveURL(tt.request.URL).Return("mock-generated-id", nil).Once()

			ts := httptest.NewServer(handlers.NewRouter(srv, cfg, logger))
			defer ts.Close()

			URL, _ := url.JoinPath(ts.URL, cfg.RoutePrefix, "api/shorten")

			var buf bytes.Buffer
			jsonEncoder := json.NewEncoder(&buf)
			_ = jsonEncoder.Encode(tt.request)

			request, _ := http.NewRequest(http.MethodPost, URL, bytes.NewReader(buf.Bytes()))
			result, _ := ts.Client().Do(request)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

			resBody, _ := io.ReadAll(result.Body)
			_ = result.Body.Close()

			buf.Reset()
			_ = json.NewEncoder(&buf).Encode(tt.want.response)
			assert.JSONEq(t, string(resBody), buf.String())
		})
	}
}
