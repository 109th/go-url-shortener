package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFHandlePost(t *testing.T) {
	type want struct {
		statusCode int
		response   string
	}
	type configs struct {
		baseURLPrefix string
		routePrefix   string
	}
	tests := []struct {
		name        string
		requestBody string
		cfg         configs
		want        want
	}{
		{
			name:        "test create redirect url",
			requestBody: "https://example.com",
			want: want{
				statusCode: 201,
				response:   "http://localhost:8080/",
			},
		},
		{
			name:        "test create redirect url with prefix",
			requestBody: "https://example.com",
			cfg: configs{
				baseURLPrefix: "http://localhost:8081/prefix",
				routePrefix:   "/prefix",
			},
			want: want{
				statusCode: 201,
				response:   "http://localhost:8081/prefix/",
			},
		},
	}

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

			mapStorage := types.NewMapStorage()
			srv := server.NewServer(mapStorage)

			ts := httptest.NewServer(handlers.Router(srv, cfg))
			defer ts.Close()

			URL, _ := url.JoinPath(ts.URL, cfg.RoutePrefix)
			request, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(tt.requestBody))
			require.NoError(t, err)
			result, err := ts.Client().Do(request)
			require.NoError(t, err)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			resBody, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			assert.Contains(t, string(resBody), tt.want.response)
			assert.Greater(t, len(string(resBody)), len(tt.want.response))
		})
	}
}
