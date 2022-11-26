package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Message       string `json:"message"`
	RealMessage   string `json:"real_message"`
	ErrorCode     int    `json:"error_code"`
	ErrorLocation string `json:"error_location"`
}

func Catch(c *fiber.Ctx) error {

	appError := recover()
	if appError != nil {
		log.Println(appError)
		return c.Status(
			appError.(AppError).ErrorCode).JSON(fiber.Map{
			"error": appError,
		})
	}

	return nil
}
