package main

import (
	"custom-modules/controller"
	"custom-modules/repository"
	"custom-modules/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// app := fiber.New()
	// userController := initializeUserController()
	// app.Get("/users", userController.FindAllUsers)
	// app.Post("/users", userController.AddUser)
	// app.Get("/users/:email", userController.FindOneByEmail)
	// app.Listen(":3000")

	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("CREATE TABLE IF NOT EXISTS users2 (id INT, name VARCHAR(255))")

}

func initializeUserController() controller.UserController {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	return controller.NewUserController(userService)
}
