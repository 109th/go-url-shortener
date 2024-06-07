package errors

import "errors"

var (
	ErrKeyExists   = errors.New("key already exists")
	ErrStorageType = errors.New("invalid storage type")
)
