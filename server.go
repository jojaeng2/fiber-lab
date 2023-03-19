package main

import (
	"custom-modules/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	controller.SetUpRoutes(app)

	app.Listen(":3000")
}
