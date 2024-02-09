package handlers

import (
	"github.com/adyssonferreira/rinha-de-backend-2024-q1/repository"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type Payload struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func ExecuteTransaction(c *fiber.Ctx) error {

	client_id := c.Params("id")

	var payload Payload
	c.BodyParser(&payload)

	client := repository.FindClientById(client_id)

	if client == nil {
		return c.SendStatus(404)
	}

	if payload.Type == "c" {
		client.Balance += payload.Value
	} else if payload.Type == "d" {

		newBalance := client.Balance - payload.Value

		if newBalance < 0 && -1*newBalance > client.Limit {
			return c.SendStatus(422)
		}

		client.Balance = newBalance

	}

	repository.UpdateBalance(client.Id, client.Balance)

	response := Response{
		Limit:   client.Limit,
		Balance: client.Balance,
	}

	return c.JSON(response)
}
