package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnectionInfo struct {
	Host     string
	Port     int32
	SSLMode  string
	DBName   string
	Username string
	Password string
}

func NewDbConnection(info ConnectionInfo) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		info.Host, info.Port, info.Username, info.Password, info.DBName, info.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
