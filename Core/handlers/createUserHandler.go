package handlers

import (
	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/infras"
	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	infras.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}
