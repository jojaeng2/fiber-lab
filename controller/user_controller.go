package controller

import (
	"custom-modules/dto"
	"custom-modules/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	AddUser(c *fiber.Ctx) error
	FindAllUsers(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService,
	}
}

func (controller *UserControllerImpl) AddUser(c *fiber.Ctx) error {
	var request dto.CreateUserRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	controller.userService.SaveUser(request)
	return c.SendStatus(fiber.StatusOK)
}

func (controller *UserControllerImpl) FindAllUsers(c *fiber.Ctx) error {
	users, err := controller.userService.FindAllUsers()
	fmt.Println("controller =", users)
	if err != nil {
		return err
	}
	return c.JSON(users)
}
