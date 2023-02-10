package main

import (
	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/infras"
	"github.com/gofiber/fiber/v2"
)

func main() {
	infras.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Xiaoxuan! DB is ready")
	})

	app.Listen(":8082")
}
