package controller

import (
	"custom-modules/dto"
	"custom-modules/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	AddUser(c *fiber.Ctx) error
	FindAllUsers(c *fiber.Ctx) error
	FindOneByEmail(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
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
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (controller *UserControllerImpl) FindOneByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := controller.userService.FindOneByEmail(email)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": email + "과 일치하는 회원이 존재하지 않습니다.",
		})
	}
	return c.JSON(user)
}

func (controller *UserControllerImpl) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	err2 := controller.userService.LoginUser(request)
	if err2 != nil {
		return err2
	}
	return c.SendStatus(fiber.StatusOK)
}
