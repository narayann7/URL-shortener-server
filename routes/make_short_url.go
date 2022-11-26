package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/models"
	srv "github.com/narayann7/gourl/services"
)

func MakeShortUrl(c *fiber.Ctx) error {
	defer srv.Catch(c)

	reqBody := new(models.Request)
	resBody := new(models.Responce)
	s := ""
	err := json.Unmarshal(c.Body(), s)
	if err != nil {

		panic(srv.AppError{
			Message:       "Unable to parse the body",
			ErrorCode:     fiber.StatusInternalServerError,
			ErrorLocation: "from MakeShortUrl, json.Unmarshal",
		})

	}
	resBody.NewUrl = reqBody.CustomShortUrl
	resBody.Url = reqBody.Url
	resBody.Expiry = reqBody.Expiry
	resBody.RateRemaining = 1

	return c.JSON(fiber.Map{
		"result": resBody,
	})

}
