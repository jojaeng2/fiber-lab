package user

import "github.com/gofiber/fiber/v2"

func SetURL(app *fiber.App) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Get /user controller!!")
	})
}
