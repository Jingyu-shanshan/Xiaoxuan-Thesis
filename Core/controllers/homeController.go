package apis

import (
	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/handlers"
	"github.com/gofiber/fiber/v2"
)

func HomeRoute(app *fiber.App) {
	app.Get("/", handlers.Home)
}
