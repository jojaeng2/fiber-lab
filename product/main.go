package main

import (
	"product/controller"
	"product/interceptor"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() != "/123" {
			return interceptor.AuthInterceptor(c)
		}
		return c.Next()
	})
	app.Use("/:id", interceptor.LoginInterceptor)
	controller := controller.NewHomeController()
	app.Get("/", controller.GetHome)
	app.Get("/:id", controller.Login)
	app.Listen(":3000")
}
