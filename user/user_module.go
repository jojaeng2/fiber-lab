package user

import (
	"custom-modules/user/controller"
	"custom-modules/user/entity"
	"custom-modules/user/repository"
	"custom-modules/user/service"

	"gorm.io/gorm"
)

type UserModule struct {
	UserRepository repository.UserRepository
	UserService    service.UserService
	UserController controller.UserController
}

func NewUserModule(db *gorm.DB) *UserModule {
	if err := db.AutoMigrate(&entity.Users{}); err != nil {
		panic("UserModule DI Fail")
	}
	repository := repository.NewUserRepository(db)
	service := service.NewUserService(repository, db)
	controller := controller.NewUserController(service)
	return &UserModule{
		repository,
		service,
		controller,
	}
}
