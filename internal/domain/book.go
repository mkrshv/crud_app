package domain

import "errors"

var (
	ErrBookNotFound = errors.New("book not found")
	ErrNotDeleted   = errors.New("book not deleted")
)

type Book struct {
	ID     int64
	Name   string
	ISBN   int64
	Author string
	Genre  string
}
