package handlers

import "github.com/adyssonferreira/rinha-de-backend-2024-q1/model"

func ExecuteTransaction() model.Tansaction {
	transaction := model.Tansaction{
		Value:       1000,
		Type:        "c",
		Description: "descicao",
	}

	return transaction
}
