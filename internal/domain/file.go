package domain

import (
	"github.com/google/uuid"
)

type File struct {
	ID  uuid.UUID
	URL string
}

func NewFile(url string) *File {
	return &File{
		ID:  uuid.New(),
		URL: url,
	}
}
