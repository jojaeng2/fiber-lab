package main

import (
	"custom-modules/user"

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
	userModule := user.NewUserModule(db)
	app.Get("/users", userModule.UserController.FindAllUsers)
	app.Post("/users", userModule.UserController.AddUser)
	app.Post("/login", userModule.UserController.Login)
	app.Get("/users/:email", userModule.UserController.FindOneByEmail)
	app.Delete("/delete/:email", userModule.UserController.DeleteByEmail)
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

func initCommentDomain(db *gorm.DB) {
}

func initBoardDomain(db *gorm.DB) {

}
