package types

import (
	"github.com/109th/go-url-shortener/internal/app/storage/errors"
)

type MapStorage struct {
	storage map[string]string
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		storage: make(map[string]string),
	}
}

func (s *MapStorage) Get(key string) string {
	return s.storage[key]
}

func (s *MapStorage) Save(key string, value string) error {
	if s.storage[key] != "" {
		return errors.ErrKeyExists
	}
	s.storage[key] = value
	return nil
}

func (s *MapStorage) Close() error {
	return nil
}
