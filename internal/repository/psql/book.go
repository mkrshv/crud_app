package psql

import (
	"database/sql"

	"github.com/mkrshv/crud_app/internal/domain"
)

//TODO: add update func

type Books struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
	return &Books{db}
}

func (b *Books) Create(book domain.Book) (int64, error) {
	res := b.db.QueryRow("INSERT INTO books (name, isbn, author, genre) VALUES ($1, $2, $3, $4) RETURNING id",
		book.Name, book.ISBN, book.Author, book.Genre)
	var id int64
	err := res.Scan(&id)
	return id, err
}

func (b *Books) GetByID(id int64) (domain.Book, error) {
	var book domain.Book
	err := b.db.QueryRow("SELECT * FROM books WHERE id=$1", id).
		Scan(&book.ID, &book.Name, &book.ISBN, &book.Author, &book.Genre)

	if err == sql.ErrNoRows {
		return book, domain.ErrBookNotFound
	}

	return book, nil
}

func (b *Books) GetAll() ([]domain.Book, error) {
	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	books := make([]domain.Book, 0)

	for rows.Next() {
		var book domain.Book
		if err = rows.Scan(&book.ID, &book.Name, &book.ISBN, &book.Author, &book.Genre); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, rows.Err()
}

func (b *Books) Delete(id int64) error {
	rows, err := b.db.Exec("DELETE FROM books WHERE id=$1", id)

	affected, _ := rows.RowsAffected()
	if affected != 1 || err != nil {
		return domain.ErrNotDeleted
	}
	return nil
}

func (b *Books) DeleteAll() {
	_, _ = b.db.Exec("DELETE FROM books")
}
