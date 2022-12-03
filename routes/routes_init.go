package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/narayann7/gourl/services"
)

func RoutesInit(app *fiber.App) {
	app.Get("/:hash", GetUrl)
	app.Get("/test/:hash", services.Test)
	app.Post("gourl/api/make-short-url", MakeShortUrl)

}
