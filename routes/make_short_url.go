package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/models"
)

func MakeShortUrl(c *fiber.Ctx) error {

	reqBody := new(models.Request)
	resBody := new(models.Responce)
	err := json.Unmarshal(c.Body(), reqBody)
	if err != nil {
		return models.AppError.SendError(
			models.AppError{
				RealMessage:   err.Error(),
				Message:       "Unable to parse the body",
				ErrorCode:     fiber.StatusInternalServerError,
				ErrorLocation: "from MakeShortUrl, json.Unmarshal",
			},
			c,
		)

	}
	resBody.NewUrl = reqBody.CustomShortUrl
	resBody.Url = reqBody.Url
	resBody.Expiry = reqBody.Expiry
	resBody.RateRemaining = 1

	return c.JSON(fiber.Map{
		"result": resBody,
	})

}
