package controller

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("user Controller!")
	})
}
