package tests

import (
	"context"
	"fmt"
	"github.com/109th/go-url-shortener/internal/app/handlers"
	"github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/mock"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	type want struct {
		statusCode     int
		response       string
		locationHeader string
	}
	tests := []struct {
		name string
		key  string
		s    *server.Server
		want want
	}{
		{
			name: "test get non existing url",
			key:  "non-existing-key",
			s:    server.NewServer(mock.NewMockStorage(map[string]string{})),
			want: want{
				statusCode: 400,
				response:   "400 bad request\n",
			},
		},
		{
			name: "test redirect url",
			key:  "ABC",
			s: server.NewServer(mock.NewMockStorage(map[string]string{
				"ABC": "https://abc.example.com",
			})),
			want: want{
				statusCode:     307,
				locationHeader: "https://abc.example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%v", tt.key), nil)

			// mock request with chi context
			rctx := chi.NewRouteContext()
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
			rctx.URLParams.Add("id", tt.key)

			w := httptest.NewRecorder()
			h := handlers.HandleGet(tt.s)
			h(w, request)

			result := w.Result()

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
