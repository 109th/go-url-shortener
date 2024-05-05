package server

import (
	"errors"
	"fmt"
	"github.com/109th/go-url-shortener/internal/app/storage"
)

var (
	ErrNotFound = errors.New("not found")
)

type Server struct {
	storage storage.IStorage
}

func NewServer(storage storage.IStorage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) GetURL(id string) (string, error) {
	value := s.storage.Get(id)
	if value == "" {
		return "", fmt.Errorf("id %s not found: %w", id, ErrNotFound)
	}

	return value, nil
}

func (s *Server) SaveURL(id string, url string) {
	s.storage.Save(id, url)
}
