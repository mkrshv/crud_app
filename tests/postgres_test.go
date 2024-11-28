package tests

import (
	"testing"

	"github.com/mkrshv/crud_app/pkg/database"
	"github.com/stretchr/testify/assert"
)

var info = database.ConnectionInfo{
	Host:     "localhost",
	Port:     5432,
	Username: "postgres",
	Password: "1234",
	DBName:   "books",
	SSLMode:  "disable",
}

func TestNewDb(t *testing.T) {

	// Убедитесь, что тестовая БД доступна перед запуском теста
	db, err := database.NewDbConnection(info)
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// Проверяем, что ошибок подключения нет
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
