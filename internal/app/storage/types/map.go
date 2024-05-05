package types

type MapStorage struct {
	storage map[string]string
}

func NewMapStorage(size int) *MapStorage {
	return &MapStorage{
		storage: make(map[string]string, size),
	}
}

func (s *MapStorage) Get(key string) string {
	return s.storage[key]
}

func (s *MapStorage) Save(key string, value string) {
	s.storage[key] = value
}
