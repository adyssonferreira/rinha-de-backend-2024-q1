package main

import (
	"github.com/adyssonferreira/rinha-de-backend-2024-q1/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//ctx := context.Background()

	app := fiber.New(fiber.Config{
		AppName:      "Rinha de Backend",
		ServerHeader: "Fiber",
	})

	app.Get("/clientes/:id/extrato", func(c *fiber.Ctx) error {
		statement := handlers.GetStatement()
		return c.JSON(statement)
	})

	app.Post("/clientes/:id/transacoes", handlers.ExecuteTransaction)

	app.Listen(":8082")
}
