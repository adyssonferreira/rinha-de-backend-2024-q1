package repository

import (
	"database/sql"
	"strconv"

	"github.com/adyssonferreira/rinha-de-backend-2024-q1/model"
	_ "github.com/lib/pq"
)

func FindClientById(id string) *model.Client {
	var client model.Client

	db, err := OpenDB()

	if err != nil {
		panic(err)
	}

	row_id, _ := strconv.Atoi(id)

	rows := db.QueryRow("SELECT * FROM clients WHERE id = $1 LIMIT 1", row_id)

	err = rows.Scan(&client.Id, &client.Nome, &client.Balance, &client.Limit)

	if err != nil {
		panic(err)
	}

	return &client
}

func UpdateBalance(client_id string, balance int) (sql.Result, error) {

	db, err := OpenDB()

	if err != nil {
		panic(err)
	}

	return db.Exec("UPDATE clients SET balance = $1 WHERE id = $2", balance, client_id)

}
