package db

import (
	"time"
)

type File struct {
	FileName  string
	FileSize  int
	Token     string
	CreatedAt time.Time
}

// BookDatabase provides thread-safe access to a database of books.
type FileDatabase interface {
	// AddBook saves a given book, assigning it a new ID.
	AddFile(b *File) (id int64, err error)

	GetFileByToken(token string) (*File, error)
}
