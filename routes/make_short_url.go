package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/models"
	srv "github.com/narayann7/gourl/services"
)

func MakeShortUrl(c *fiber.Ctx) error {
	//recover the panic and send as respond
	//with suitable message and status code
	defer srv.CatchErrors(c)
	//create new request Struct
	reqBody := new(models.Request)
	//parce the json string to Request Struct
	err := json.Unmarshal(c.Body(), reqBody)
	//if error painc
	if err != nil {
		panic(srv.AppError{
			RealMessage:   err.Error(),
			Message:       "Unable to parse the body",
			ErrorCode:     fiber.StatusInternalServerError,
			ErrorLocation: "from MakeShortUrl, json.Unmarshal",
		})
	}
	//rate limiting

	//	check for valid url
	if !srv.IsVaildUrl(reqBody.Url) {
		errorMessage := ""
		if reqBody.Url == "" {
			errorMessage = "url parameter is require"
		} else {
			errorMessage = "not a valid url"
		}
		panic(srv.AppError{
			Message:   errorMessage,
			ErrorCode: 400,
		})

	}

	//check for domain error
	if err := srv.CheckForDominError(reqBody.Url); !err {
		panic(srv.AppError{
			Message:   "this url is not allowed",
			ErrorCode: 400,
		})
	}

	//enforce http
	reqBody.Url = srv.EnforceHttp(reqBody.Url)

	srv.GetNewHash()

	return c.JSON(fiber.Map{
		"result": reqBody,
	})

}
