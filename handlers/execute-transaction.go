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

const DEBIT_TRANSACTION = "d"
const CREDIT_TRANSACTION = "c"

func ExecuteTransaction(c *fiber.Ctx) error {

	client_id := c.Params("id")

	var payload Payload
	c.BodyParser(&payload)

	client, err := repository.FindClientById(client_id)

	// Client not found
	if client == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound, "Client not found")
	}

	// Internal error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	if payload.Type == CREDIT_TRANSACTION {
		client.Balance += payload.Value
	} else if payload.Type == DEBIT_TRANSACTION {

		newBalance := client.Balance - payload.Value

		if newBalance < 0 && -1*newBalance > client.Limit {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Unprocessable Entity")
		}

		client.Balance = newBalance

	}

	_, err = repository.UpdateBalance(client.Id, client.Balance)

	// Internal error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	response := Response{
		Limit:   client.Limit,
		Balance: client.Balance,
	}

	return c.JSON(response)
}
