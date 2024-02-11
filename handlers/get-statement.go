package handlers

import (
	"time"

	"github.com/adyssonferreira/rinha-de-backend-2024-q1/model"
	"github.com/adyssonferreira/rinha-de-backend-2024-q1/repository"
	"github.com/gofiber/fiber/v2"
)

func GetStatement(c *fiber.Ctx) error {

	client_id := c.Params("id")

	client, err := repository.FindClientById(client_id)

	// Client not found
	if client == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound, "Client not found")
	}

	transactions, err := repository.FindTransactionsByClientId(client_id)

	// Internal error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error on find transations")
	}

	statement := model.Statement{
		Balance: model.Balance{
			Date:  time.Now().Format(time.RFC3339),
			Total: client.Balance,
			Limit: client.Limit,
		},

		Tansactions: transactions,
	}

	return c.JSON(statement)
}
