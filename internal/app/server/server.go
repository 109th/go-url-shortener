package server

import (
	"errors"
	"fmt"

	"github.com/109th/go-url-shortener/internal/app/storage"
)

const RandomBytesCount = 8

var (
	ErrNotFound = errors.New("not found")
)

type Server struct {
	storage Storage
}

type Storage interface {
	Get(key string) string
	Save(key, value string) error
}

func NewServer(s Storage) *Server {
	return &Server{
		storage: s,
	}
}

func (s *Server) GetURL(id string) (string, error) {
	value := s.storage.Get(id)
	if value == "" {
		return "", fmt.Errorf("id %s not found: %w", id, ErrNotFound)
	}

	return value, nil
}

func (s *Server) SaveURL(url string) (string, error) {
	var uid string
	var err error

	// max 10 attempts to store the URL
	for range 10 {
		uid, err = GetRandomString(RandomBytesCount)
		if err != nil {
			return "", err
		}

		err = s.storage.Save(uid, url)
		if errors.Is(err, storage.ErrKeyExists) {
			continue
		}

		if err == nil {
			return uid, nil
		}
	}

	return "", fmt.Errorf("can't save URL: %w", err)
}
