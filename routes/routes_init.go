package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/services"
)

func RoutesInit(app *fiber.App) {
	app.Get("/:hash", GetUrl)
	app.Post("gourl/api/make-short-url", MakeShortUrl)
	app.Get("gourl/api/generate-customs-urls/:size", GenerateNewHashes)

	//------------testing---------------
	app.Get("/test/:hash", services.Test)
}
