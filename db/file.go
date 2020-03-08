package db

import (
  "time"
)

type File struct {
  FileName          string
  Token             string
  CreatedAt         time.Time
  UpdatedAt         time.Time
}

// BookDatabase provides thread-safe access to a database of books.
type FileDatabase interface {
	// AddBook saves a given book, assigning it a new ID.
	AddFile(b *File) (id int64, err error)

  GetFileByToken(token string) (*File, error)
}


