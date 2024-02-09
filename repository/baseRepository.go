package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil

func OpenDB() (*sql.DB, error) {

	if db == nil {
		db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=pass dbname=postgres sslmode=disable")
		return db, err
	}

	return db, nil
}
