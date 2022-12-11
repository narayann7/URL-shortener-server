package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/services"
)

func RoutesInit(app *fiber.App) {
	//------------testing---------------
	app.Get("/test", services.Test)
	app.Get("/:hash", GetUrl)
	app.Get("gourl/api/init", BrowserInit)
	app.Post("gourl/api/make-short-url", MakeShortUrl)
	app.Get("gourl/api/generate-customs-urls/:size", GenerateNewHashes)

}
