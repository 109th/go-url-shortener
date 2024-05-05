package handlers

import (
	srv "github.com/109th/go-url-shortener/internal/app/server"
	"github.com/109th/go-url-shortener/internal/app/storage/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePost(t *testing.T) {
	type want struct {
		statusCode int
		response   string
	}
	tests := []struct {
		name        string
		requestBody string
		s           *srv.Server
		want        want
	}{
		{
			name:        "test create redirect url",
			requestBody: "https://example.com",
			s:           srv.NewServer(mock.NewMockStorage(map[string]string{})),
			want: want{
				statusCode: 201,
				response:   "http://localhost:8080/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()
			h := HandlePost(tt.s)
			h(w, request)

			result := w.Result()

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
