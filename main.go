package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/narayann7/gourl/routes"
	"github.com/narayann7/gourl/services"
)

func main() {

	if isDev, _ := strconv.ParseBool(services.Getenv("IS_DEV")); isDev {
		app := fiber.New()
		//-----------init--------------//
		//logger for logging request
		app.Use(logger.New())
		//routes initiliztion
		routes.RoutesInit(app)
		//port from env
		port := ":" + services.Getenv("PORT")
		listenError := app.Listen(port)
		if listenError != nil {
			log.Fatalf("something went wrong %v", listenError)
		}
	} else {
		services.Testing()
	}
}
