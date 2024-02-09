package main

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type Usuario struct {
	/* Testgroup int    `json:"testgroup"  redis:"testgroup"`
	Status    string `json:"status" redis:"status"` */

	Nome  string
	Idade int
}

func main() {

	ctx := context.Background()

	app := fiber.New(fiber.Config{
		AppName:      "Rinha de Backend",
		ServerHeader: "Fiber",
	})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "pass", // no password set
		DB:       0,      // use default DB
	})

	datasave, _ := json.Marshal(Usuario{Nome: "José", Idade: 12})

	rdb.Set(ctx, "user", datasave, 0)

	app.Get("/clientes/:id/extrato", func(c fiber.Ctx) error {

		var user Usuario

		result := rdb.Get(ctx, "user")
		json.Unmarshal([]byte(result.Val()), &user)

		return c.JSON(user)
	})

	app.Post("/clientes/:id/extrato", func(c fiber.Ctx) error {
		return c.SendString("I'm a Post request! Param:" + c.Params("id"))
	})

	app.Listen(":8082")
}
