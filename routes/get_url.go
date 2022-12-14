package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/database"
	"github.com/narayann7/gourl/services"
)

func GetUrl(c *fiber.Ctx) error {

	rdb0 := database.DatabaseInit(0)
	defer rdb0.Close()

	//rate limiting
	err := database.CounterReducer(c.IP())
	if err != nil {
		panic(services.AppError{
			Message:   err.Error(),
			ErrorCode: 400,
		})
	}

	data, err := rdb0.Get(database.Ctx, c.Params("hash")).Result()
	//if err is not nil which means the hash has no value to it
	//and is not used for any url
	//length of data == 0 means empty string came
	if len(data) == 0 && err != nil {
		return c.
			JSON(fiber.Map{
				"url": err,
			})
	} else {
		return c.Status(307).Redirect(data)
	}

}
