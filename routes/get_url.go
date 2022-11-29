package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/database"
)

func GetUrl(c *fiber.Ctx) error {

	rdb0 := database.DatabaseInit(0)

	rdb0.Set(database.Ctx, c.Params("hash"), "hello.com", time.Second*50000000)
	return c.JSON(fiber.Map{
		"type": c.Params("hash"),
	})
}
