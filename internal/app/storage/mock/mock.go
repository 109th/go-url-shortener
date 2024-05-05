package mock

type Storage struct {
	storage map[string]string
}

func NewMockStorage(values map[string]string) *Storage {
	return &Storage{
		storage: values,
	}
}

func (s *Storage) Get(key string) string {
	return s.storage[key]
}

func (s *Storage) Save(key, value string) {
	s.storage[key] = value
}
