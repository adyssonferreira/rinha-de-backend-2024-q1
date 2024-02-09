package main

import (
	"github.com/adyssonferreira/rinha-de-backend-2024-q1/handlers"
	"github.com/gofiber/fiber/v3"
)

func main() {

	//ctx := context.Background()

	app := fiber.New(fiber.Config{
		AppName:      "Rinha de Backend",
		ServerHeader: "Fiber",
	})

	app.Get("/clientes/:id/extrato", func(c fiber.Ctx) error {
		statement := handlers.GetStatement()
		return c.JSON(statement)
	})

	app.Post("/clientes/:id/transacoes", func(c fiber.Ctx) error {
		transaction := handlers.ExecuteTransaction()
		return c.JSON(transaction)
	})

	app.Listen(":8082")
}
