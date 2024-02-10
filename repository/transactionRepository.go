package repository

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
