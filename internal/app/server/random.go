package server

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GetRandomString(size int) (string, error) {
	data := make([]byte, size)
	if _, err := rand.Read(data); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return base64.RawURLEncoding.EncodeToString(data), nil
}
