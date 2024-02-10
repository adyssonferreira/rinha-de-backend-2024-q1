package handlers

import (
	"github.com/adyssonferreira/rinha-de-backend-2024-q1/constants"
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

	if len(payload.Description) > 10 {
		return fiber.NewError(fiber.StatusBadRequest, "Description to long")
	}

	client, err := repository.FindClientById(client_id)

	// Client not found
	if client == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound, "Client not found")
	}

	// Internal error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	if payload.Type == constants.CREDIT_TRANSACTION {
		client.Balance += payload.Value
	} else if payload.Type == constants.DEBIT_TRANSACTION {

		newBalance := client.Balance - payload.Value

		if newBalance < 0 && -1*newBalance > client.Limit {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Unprocessable Entity")
		}

		client.Balance = newBalance

	}

	err = repository.CreateTransaction(client.Id, payload.Value, payload.Type, payload.Description)

	// Internal error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error on create transation")
	}

	err = repository.UpdateBalance(client.Id, client.Balance)

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
