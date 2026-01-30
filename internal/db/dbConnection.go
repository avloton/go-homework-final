package db

import (
	"database/sql"
	"log"
	_"modernc.org/sqlite"
)

func DbConnect() *sql.DB {
	db, err := sql.Open("sqlite", "./db/bakery.db")
	if err != nil {
		log.Fatalf("Couldn't connect to db: %v\n", err)
	}
	return db
}