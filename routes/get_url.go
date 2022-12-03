package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/database"
)

func GetUrl(c *fiber.Ctx) error {

	rdb0 := database.DatabaseInit(0)

	dd, ee := rdb0.Set(database.Ctx, c.Params("hash"), "hello.com", time.Second*50000000).Result()
	fmt.Println(dd)
	fmt.Println(ee)
	return c.JSON(fiber.Map{
		"type": c.Params("hash"),
	})
}
