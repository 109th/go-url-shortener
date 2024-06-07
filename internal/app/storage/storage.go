package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/109th/go-url-shortener/internal/app/config"
	"github.com/109th/go-url-shortener/internal/app/storage/errors"
	"github.com/109th/go-url-shortener/internal/app/storage/types"
)

type Storage interface {
	Get(key string) string
	Save(key, value string) error
	Close() error
}

func NewStorage(cfg *config.Config) (Storage, error) {
	switch cfg.StorageType {
	case "memory":
		return types.NewMapStorage(), nil
	case "file":
		//nolint:gomnd // magic allowed here ;)
		file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			log.Fatalf("open file error: %v", err)
		}

		storage, err := types.NewFileStorage(file)
		if err != nil {
			return nil, fmt.Errorf("create storage error: %w", err)
		}

		return storage, nil
	}

	return nil, errors.ErrStorageType
}
