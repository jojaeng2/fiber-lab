package main

import (
	"custom-modules/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	user.SetURL(app)
	app.Listen(":3000")
}
