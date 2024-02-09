package handlers

import "github.com/adyssonferreira/rinha-de-backend-2024-q1/model"

func GetStatement() model.Statement {
	statement := model.Statement{
		Balance: model.Balance{
			Total: 1220,
			Date:  "2024-01-17T02:34:41.217753Z",
			Limit: 100000,
		},

		Tansactions: []model.Tansaction{
			{
				Value:       10,
				Type:        "C",
				Description: "descricao",
				CreateAt:    "2024-01-17T02:34:38.543030Z",
			},
			{
				Value:       20,
				Type:        "C",
				Description: "descricao 2",
				CreateAt:    "2024-01-17T02:34:38.543030Z",
			},
		},
	}

	return statement

}
