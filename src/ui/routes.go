package ui

import (
	"github.com/gofiber/fiber/v3"
)

func (cb *controllersImpl) InitRoutes(app *fiber.App) {
	app.Get("/user/login", func(c fiber.Ctx) error {
		cb.GetUserController().GetAppService().LoginBySocial()

		return c.JSON("")
	})
	app.Post("/session", func(c fiber.Ctx) error {

		return c.JSON("Store")
	})
	app.Get("/session", func(c fiber.Ctx) error {
		//cb.GetUserController().GetAppService().LoginBySocial()

		return c.JSON("is logged")
	})
}
