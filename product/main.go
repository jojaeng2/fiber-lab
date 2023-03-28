package main

import (
	"fmt"
	"product/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	controller := controller.NewHomeController()
	app.Get("/", controller.GetHome)
	fmt.Print("Hello")
	app.Listen(":3000")

}
