package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFHandleGet(t *testing.T) {
	type want struct {
		statusCode     int
		response       string
		locationHeader string
	}
	type configs struct {
		routePrefix string
	}
	tests := []struct {
		name string
		key  string
		data map[string]string
		cfg  configs
		want want
	}{
		{
			name: "test get non existing url",
			key:  "non-existing-key",
			data: nil,
			want: want{
				statusCode: 400,
				response:   "400 bad request\n",
			},
		},
		{
			name: "test redirect url",
			key:  "ABC",
			data: map[string]string{
				"ABC": "https://abc.example.com",
			},
			want: want{
				statusCode:     307,
				locationHeader: "https://abc.example.com",
			},
		},
		{
			name: "test get non existing url with prefix",
			key:  "non-existing-key",
			data: nil,
			cfg: configs{
				routePrefix: "/prefix",
			},
			want: want{
				statusCode: 400,
				response:   "400 bad request\n",
			},
		},
		{
			name: "test redirect url with prefix",
			key:  "ABC",
			data: map[string]string{
				"ABC": "https://abc.example.com",
			},
			cfg: configs{
				routePrefix: "/prefix",
			},
			want: want{
				statusCode:     307,
				locationHeader: "https://abc.example.com",
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

			// restore default config
			defer func() {
				config.RoutePrefix = config.DefaultRoutePrefix
			}()

			mapStorage := types.NewMapStorage()
			srv := server.NewServer(mapStorage)
			for key, value := range tt.data {
				err := mapStorage.Save(key, value)
				require.NoError(t, err)
			}

			ts := httptest.NewServer(handlers.Router(srv, config.RoutePrefix))
			ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
			defer ts.Close()

			URL := ts.URL + strings.TrimRight(config.RoutePrefix, "/")
			request, err := http.NewRequest(http.MethodGet, URL+fmt.Sprintf("/%v", tt.key), nil)
			require.NoError(t, err)
			result, err := ts.Client().Do(request)
			require.NoError(t, err)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			if tt.want.response != "" {
				resBody, err := io.ReadAll(result.Body)
				require.NoError(t, err)
				err = result.Body.Close()
				require.NoError(t, err)

				assert.Equal(t, tt.want.response, string(resBody))
			}

			if tt.want.locationHeader != "" {
				assert.Equal(t, result.Header.Get("Location"), tt.want.locationHeader)
			}
		})
	}
}
