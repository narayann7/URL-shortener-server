package services

import (
	"log"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Message       string `json:"message"`
	RealMessage   string `json:"real_message"`
	ErrorCode     int    `json:"error_code"`
	ErrorLocation string `json:"error_location"`
}

/*
general panic code is must
	panic(fiber.Map{
			"code":    400,
			"message": "testing from map",
		})

*/

func Catch(c *fiber.Ctx) error {

	appError := recover()
	if appError != nil {
		// log.Println(appError)

		var interFacetype = (reflect.TypeOf(appError))
		log.Printf("interFace type ------> " + interFacetype.String())
		if interFacetype.String() == "services.AppError" {

			if appError.(AppError).RealMessage == "" {
				return c.Status(
					appError.(AppError).ErrorCode).JSON(fiber.Map{
					"error": appError.(AppError).Message,
				})
			} else {
				return c.Status(
					appError.(AppError).ErrorCode).JSON(fiber.Map{
					"error": appError,
				})
			}

		} else if interFacetype.String() == "string" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": appError,
			})
		} else if interFacetype.String() == "fiber.Map" {
			temp := appError.(fiber.Map)
			code := temp["code"]

			return c.Status(code.(int)).JSON(fiber.Map{
				"error": temp,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Something went wrong",
			})
		}

	}

	return nil
}
