package main

import (
	"custom-modules/user/controller"
	userentity "custom-modules/user/entity"
	"custom-modules/user/repository"
	"custom-modules/user/service"

	// "custom-modules/comment/controller"
	// boardentity "custom-modules/comment/entity"
	// "custom-modules/comment/repository"
	// "custom-modules/comment/service"

	// "custom-modules/board/controller"
	// "custom-modules/board/entity"
	// "custom-modules/board/repository"
	// "custom-modules/board/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	db := initDB()
	userController := initUserDomain(db)
	app.Get("/users", userController.FindAllUsers)
	app.Post("/users", userController.AddUser)
	app.Post("/login", userController.Login)
	app.Get("/users/:email", userController.FindOneByEmail)
	app.Delete("/delete/:email", userController.DeleteByEmail)
	app.Listen(":3000")
}

func initDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func initUserDomain(db *gorm.DB) controller.UserController {
	db.AutoMigrate(&userentity.Users{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db)
	return controller.NewUserController(userService)
}

func initCommentDomain(db *gorm.DB) {
}

func initBoardDomain(db *gorm.DB) {

}
