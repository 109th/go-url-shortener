package tests

import (
	"github.com/109th/go-url-shortener/internal/app/handlers"
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
	tests := []struct {
		name        string
		requestBody string
		s           *server.Server
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(handlers.Router(tt.s))
			defer ts.Close()

			request, err := http.NewRequest(http.MethodPost, ts.URL, strings.NewReader(tt.requestBody))
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
