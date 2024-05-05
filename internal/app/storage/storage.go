package storage

type IStorage interface {
	Get(key string) string
	Save(key, value string)
}
