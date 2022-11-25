package routes

import "github.com/gofiber/fiber/v2"

func RoutesInit(app *fiber.App) {
	app.Get("gourl/api/get-url", GetUrl)
	app.Post("gourl/api/make-short-url", MakeShortUrl)

}
