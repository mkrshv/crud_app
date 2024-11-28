package main

import (
	"log"

	"github.com/mkrshv/crud_app/pkg/database"
)

func main() {
	connectionInfo := database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "1234",
		DBName:   "books",
		SSLMode:  "disable",
	}

	db, err := database.NewDbConnection(connectionInfo)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
