package main

import (
	"custom-modules/controller"
	"custom-modules/entity"
	"custom-modules/repository"
	"custom-modules/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// User 모델에 해당하는 테이블 생성
	db.AutoMigrate(&entity.Users{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	return controller.NewUserController(userService)
}
