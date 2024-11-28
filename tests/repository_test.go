package tests

import (
	"testing"

	"github.com/mkrshv/crud_app/internal/domain"
	"github.com/mkrshv/crud_app/internal/repository/psql"
	"github.com/mkrshv/crud_app/pkg/database"
	"github.com/stretchr/testify/assert"
)

//var books []domain.Book

var books = []domain.Book{
	{
		Name:   "Война и мир",
		ISBN:   1245331212,
		Author: "Л.Н. Толстой",
		Genre:  "Роман",
	},
	{
		Name:   "Герой нашего времени",
		ISBN:   124325543,
		Author: "М.Ю. Лермонтов",
		Genre:  "Роман",
	},
}

func TestRepo(t *testing.T) {
	db, _ := database.NewDbConnection(info)
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	repo := psql.NewBooks(db)
	defer repo.DeleteAll()

	id, err := repo.Create(books[0])
	assert.NoError(t, err)

	book, err := repo.GetByID(id)
	assert.NoError(t, err)

	_, err = repo.Create(books[1])
	assert.NoError(t, err)
	books, err := repo.GetAll()
	assert.NoError(t, err)
	t.Log(books)

	t.Log(book)

}
