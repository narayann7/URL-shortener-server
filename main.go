package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/routes"
	"github.com/narayann7/gourl/services"
)

func main() {

	app := fiber.New()

	routes.RoutesInit(app)

	port := ":" + services.Getenv("PORT")
	listenError := app.Listen(port)
	if listenError != nil {
		log.Fatalf("something went wrong %v", listenError)
	}
}
