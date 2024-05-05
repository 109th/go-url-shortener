package storage

var storage map[string]string

func Init(size int) {
	if storage == nil {
		storage = make(map[string]string, size)
	}
}

func Get(key string) string {
	return storage[key]
}

func Set(key, value string) {
	storage[key] = value
}
