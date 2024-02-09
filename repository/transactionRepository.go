package repository

import (
	"database/sql"
	"errors"
)

func CreateTransaction(client_id int, value int, transactionType string, description string) (sql.Result, error) {
	db, err := OpenDB()

	if err != nil {
		return nil, errors.New("can't open database")
	}

	return db.Exec("INSERT INTO transactions(client_id, value, type, description) VALUES ($1,$2,$3,$4)", client_id, value, transactionType, description)

}
