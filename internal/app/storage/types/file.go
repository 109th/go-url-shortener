package types

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/109th/go-url-shortener/internal/app/storage/errors"
)

type FileStorage struct {
	file *os.File
	*MapStorage
}

type storageRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewFileStorage(file *os.File) (*FileStorage, error) {
	s := &FileStorage{
		file:       file,
		MapStorage: NewMapStorage(),
	}

	err := fillMapStorage(s)
	if err != nil {
		return nil, fmt.Errorf("fill map storage error: %w", err)
	}

	return s, nil
}

func (s *FileStorage) Get(key string) string {
	return s.MapStorage.Get(key)
}

func (s *FileStorage) Save(key string, value string) error {
	if s.MapStorage.Get(key) != "" {
		return errors.ErrKeyExists
	}

	err := s.MapStorage.Save(key, value)
	if err != nil {
		return err
	}

	record := &storageRecord{Key: key, Value: value}
	data, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("json encode error: %w", err)
	}
	data = append(data, '\n')

	_, err = s.file.Write(data)
	if err != nil {
		return fmt.Errorf("write to file error: %w", err)
	}

	return nil
}

func (s *FileStorage) Close() error {
	err := s.file.Close()
	if err != nil {
		return fmt.Errorf("close file storage error: %w", err)
	}

	return nil
}

func fillMapStorage(s *FileStorage) error {
	reader := bufio.NewScanner(s.file)
	for reader.Scan() {
		token := reader.Bytes()
		data := &storageRecord{}

		err := json.Unmarshal(token, data)
		if err != nil {
			return fmt.Errorf("json decode error: %w", err)
		}

		err = s.MapStorage.Save(data.Key, data.Value)
		if err != nil {
			return fmt.Errorf("save error: %w", err)
		}
	}

	return nil
}
