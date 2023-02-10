package apis

import (
	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", handlers.Home)
}
