package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/services"
)

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"hello": "world",
			},
		)
	})

	port := ":" + services.Getenv("PORT")
	listenError := app.Listen(port)
	if listenError != nil {
		log.Fatalf("something went wrong %v", listenError)
	}
}
