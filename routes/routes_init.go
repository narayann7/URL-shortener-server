package routes

import "github.com/gofiber/fiber/v2"

func RoutesInit(app *fiber.App) {
	app.Get("/:hash", GetUrl)
	app.Post("gourl/api/make-short-url", MakeShortUrl)

}
