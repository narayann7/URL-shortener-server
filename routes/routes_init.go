package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {

	app.Get("/:hash", GetUrl)
	app.Get("gourl/api/init", BrowserInit)
	app.Post("gourl/api/make-short-url", MakeShortUrl)
	app.Get("gourl/api/generate-customs-urls/:size", GenerateNewHashes)

}
