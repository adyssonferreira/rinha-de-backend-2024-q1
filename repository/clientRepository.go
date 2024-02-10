package repository

import (
	"errors"

	"github.com/adyssonferreira/rinha-de-backend-2024-q1/model"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func FindClientById(id string) (*model.Client, error) {
	var client model.Client

	db, err := OpenDB()

	if err != nil {
		return nil, errors.New("can't open database")
	}

	err = db.QueryRow(contextt, "SELECT * FROM clients WHERE id = $1 LIMIT 1", id).Scan(&client.Id, &client.Nome, &client.Balance, &client.Limit)

	// NOT FOUND
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &client, nil
}

func UpdateBalance(client_id int, balance int) error {

	db, err := OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec(contextt, "UPDATE clients SET balance = $1 WHERE id = $2", balance, client_id)
	if err != nil {
		return err
	}

	return nil
}
