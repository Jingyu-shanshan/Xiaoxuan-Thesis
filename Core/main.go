package main

import (
	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/infras"
	"github.com/gofiber/fiber/v2"
)

func main() {
	infras.ConnectDb()

	app := fiber.New()

	infras.setupRoutes(app)

	app.Listen(":8082")
}
