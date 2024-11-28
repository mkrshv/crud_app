package service

import "github.com/mkrshv/crud_app/internal/domain"

type BooksRepository interface {
	Create(book domain.Book) (int64, error)
	GetByID(id int64) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int64) error
}

type Books struct {
	repo BooksRepository
}

func NewBooks(repo BooksRepository) *Books {
	return &Books{repo: repo}
}

func (b *Books) Create(book domain.Book) (int64, error) {
	return b.repo.Create(book)
}

func (b *Books) GetByID(id int64) (domain.Book, error) {
	return b.repo.GetByID(id)
}

func (b *Books) GetAll() ([]domain.Book, error) {
	return b.repo.GetAll()
}

func (b *Books) Delete(id int64) error {
	return b.repo.Delete(id)
}
