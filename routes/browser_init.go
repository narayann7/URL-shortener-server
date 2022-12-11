package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/database"
	"github.com/narayann7/gourl/services"
)

func BrowserInit(c *fiber.Ctx) error {
	defer services.CatchErrors(c)
	rdb1 := database.DatabaseInit(1)
	defer rdb1.Close()

	limit, err := rdb1.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		quota, _ := strconv.Atoi(os.Getenv("API_QUOTA"))
		status, err := rdb1.Set(database.Ctx, c.IP(), quota, time.Minute*10).Result()
		if status == "OK" && err == nil {
			return c.JSON(fiber.Map{
				"ip_address": c.IP(),
				"api_quota":  quota,
			})
		} else {
			panic(services.AppError{
				Message:   "Something went wrong",
				ErrorCode: 500,
			})
		}
	} else {
		quota, _ := strconv.Atoi(limit)

		return c.JSON(fiber.Map{
			"ip_address": c.IP(),
			"api_quota":  quota,
		})
	}

}
