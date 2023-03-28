package main

import (
	"fmt"
	"product/controller"
	"product/interceptor"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use("/", interceptor.AuthInterceptor)
	app.Use("/:id", interceptor.LoginInterceptor)
	controller := controller.NewHomeController()
	app.Get("/", controller.GetHome)
	app.Get("/:id", controller.Login)
	fmt.Print("Hello")
	app.Listen(":3000")

}
