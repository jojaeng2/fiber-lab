package main

import (
	"custom-modules/controller"
	"custom-modules/repository"
	"custom-modules/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	userController := initializeUserController()
	app.Get("/users", userController.FindAllUsers)
	app.Post("/users", userController.AddUser)
	app.Get("/users/:email", userController.FindOneByEmail)
	app.Listen(":3000")
}

func initializeUserController() controller.UserController {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	return controller.NewUserController(userService)
}
