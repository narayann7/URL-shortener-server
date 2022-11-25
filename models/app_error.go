package models

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Message       string `json:"message"`
	RealMessage   string `json:"real_message"`
	ErrorCode     int    `json:"error_code"`
	ErrorLocation string `json:"error_location"`
}

func (appError AppError) SendError(c *fiber.Ctx) error {

	return c.Status(appError.ErrorCode).JSON(fiber.Map{
		"error": appError,
	})

}
