package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetUrl(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"type": "GetUrl",
	})
}
