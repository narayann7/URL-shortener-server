package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/models"
)

func MakeShortUrl(c *fiber.Ctx) error {

	reqBody := new(models.Request)
	resBody := new(models.Responce)
	ss := "12"
	err := json.Unmarshal(c.Body(), ss)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	resBody.NewUrl = reqBody.CustomShortUrl
	resBody.Url = reqBody.Url
	resBody.Expiry = reqBody.Expiry
	resBody.RateRemaining = 1

	return c.JSON(fiber.Map{
		"result": resBody,
	})

}
