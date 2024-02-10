package cache

import (
	"github.com/adyssonferreira/rinha-de-backend-2024-q1/model"
)

var ClientStore = map[string]model.Client{}

func SaveUser(client_id string, client model.Client) {
	ClientStore[client_id] = client
}

func GetUser(client_id string) *model.Client {
	client, exists := ClientStore[client_id]

	if exists {
		return &client
	}

	return &model.Client{}
}
