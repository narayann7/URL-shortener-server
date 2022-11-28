package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/models"
	"github.com/narayann7/gourl/services"
	srv "github.com/narayann7/gourl/services"
)

func MakeShortUrl(c *fiber.Ctx) error {
	defer srv.CatchErrors(c)

	reqBody := new(models.Request)
	err := json.Unmarshal(c.Body(), reqBody)
	if err != nil {
		panic(srv.AppError{
			RealMessage:   err.Error(),
			Message:       "Unable to parse the body",
			ErrorCode:     fiber.StatusInternalServerError,
			ErrorLocation: "from MakeShortUrl, json.Unmarshal",
		})
	}
	//rate limiting

	//check for valid url
	if !services.IsVaildUrl(reqBody.Url) {
		panic(srv.AppError{
			Message:   "not a valid url",
			ErrorCode: 400,
		})

	}

	//check for domain error

	//enforce http

	return c.JSON(fiber.Map{
		"result": reqBody,
	})

}
