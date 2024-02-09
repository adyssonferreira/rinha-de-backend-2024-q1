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

	/* rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "pass", // no password set
		DB:       0,      // use default DB
	}) */

	//rdb.Set(ctx, "user", datasave, 0)

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
