package domain

import (
	"github.com/google/uuid"
)

type File struct {
	ID  uuid.UUID
	URL string
	Key string
}

func NewFile(url, key string) *File {
	return &File{
		ID:  uuid.New(),
		URL: url,
		Key: key,
	}
}
