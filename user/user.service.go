package user

import "github.com/gofiber/fiber/v2"

type UserService struct {
	name string
}

func GetUser(ctx *fiber.Ctx) string {
	userService := NewUserService()
	return userService.name
}

func NewUserService() UserService {
	return UserService{"HO"}
}
