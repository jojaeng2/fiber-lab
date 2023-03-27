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
	db := initDB()
	userController := initUserController(db)
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

func initUserController(db *gorm.DB) controller.UserController {
	if err := db.AutoMigrate(&entity.Users{}); err != nil {
		panic("UserModule DI Fail")
	}
	repository := repository.NewUserRepository(db)
	service := service.NewUserService(repository, db)
	controller := controller.NewUserController(service)
	return controller
}

func initCommentDomain(db *gorm.DB) {
}

func initBoardDomain(db *gorm.DB) {

}
