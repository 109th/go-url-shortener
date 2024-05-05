package tests

import (
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/handlers/config"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
		s           *server.Server
		cfg         configs
		want        want
	}{
		{
			name:        "test create redirect url",
			requestBody: "https://example.com",
			s:           server.NewServer(mock.NewMockStorage(map[string]string{})),
			want: want{
				statusCode: 201,
				response:   "http://localhost:8080/",
			},
		},
		{
			name:        "test create redirect url with prefix",
			requestBody: "https://example.com",
			s:           server.NewServer(mock.NewMockStorage(map[string]string{})),
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
			config.RoutePrefix = config.DefaultRoutePrefix
			if tt.cfg.routePrefix != "" {
				config.RoutePrefix = tt.cfg.routePrefix
			}

			config.ServerURLPrefix = config.DefaultServerURL
			if tt.cfg.baseURLPrefix != "" {
				config.ServerURLPrefix = tt.cfg.baseURLPrefix
			}

			// restore default config
			defer func() {
				config.RoutePrefix = config.DefaultRoutePrefix
				config.ServerURLPrefix = config.DefaultServerURL
			}()

			ts := httptest.NewServer(handlers.Router(tt.s, config.RoutePrefix))
			defer ts.Close()

			URL := ts.URL + strings.TrimRight(config.RoutePrefix, "/")
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
