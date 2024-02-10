package repository

import "github.com/adyssonferreira/rinha-de-backend-2024-q1/model"

func CreateTransaction(client_id int, value int, transactionType string, description string) error {

	db, err := OpenDB()

	if err != nil {
		return err
	}

	_, err = db.Exec(contextt, "INSERT INTO transactions(client_id, value, type, description) VALUES ($1,$2,$3,$4)", client_id, value, transactionType, description)
	if err != nil {
		return err
	}

	return nil
}

func FindTransactionsByClientId(client_id string) ([]model.Tansaction, error) {

	db, err := OpenDB()

	if err != nil {
		return []model.Tansaction{}, err
	}

	transactions := []model.Tansaction{}

	rows, err := db.Query(contextt, "SELECT value, type, description, create_at FROM transactions WHERE client_id = $1", client_id)

	if err != nil {
		return []model.Tansaction{}, err
	}

	for rows.Next() {
		var transaction model.Tansaction
		rows.Scan(&transaction.Value, &transaction.Type, &transaction.Description, &transaction.CreateAt)

		transactions = append(transactions, transaction)
	}

	return transactions, nil

}
