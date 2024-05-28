package types

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/109th/go-url-shortener/internal/app/storage"
)

type MapStorage struct {
	file    *os.File
	storage map[string]string
}

type storageRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewMapStorage(file *os.File) (*MapStorage, error) {
	s := &MapStorage{
		file:    file,
		storage: make(map[string]string),
	}

	err := fillMapStorage(s)
	if err != nil {
		return nil, fmt.Errorf("fill map storage error: %w", err)
	}

	return s, nil
}

func fillMapStorage(m *MapStorage) error {
	reader := bufio.NewScanner(m.file)
	for reader.Scan() {
		token := reader.Bytes()
		data := &storageRecord{}

		err := json.Unmarshal(token, data)
		if err != nil {
			return fmt.Errorf("json decode error: %w", err)
		}

		m.storage[data.Key] = data.Value
	}

	return nil
}

func (s *MapStorage) Get(key string) string {
	return s.storage[key]
}

func (s *MapStorage) Save(key string, value string) error {
	if s.storage[key] != "" {
		return storage.ErrKeyExists
	}
	s.storage[key] = value

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
