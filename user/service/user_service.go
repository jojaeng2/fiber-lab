package service

import (
	"custom-modules/user/dto"
	"custom-modules/user/entity"
	"custom-modules/user/repository"
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	SaveUser(request dto.CreateUserRequest) error
	FindAllUsers() ([]entity.Users, error)
	FindOneByEmail(email string) (interface{}, error)
	LoginUser(request dto.LoginRequest) error
	DeleteUserByEmail(email string) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		db:             db,
	}
}

func (userService *UserServiceImpl) SaveUser(request dto.CreateUserRequest) error {
	return userService.db.Transaction(func(tx *gorm.DB) error {
		user := entity.Users{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		}
		return userService.userRepository.Save(&user)
	})
}

func (userService *UserServiceImpl) FindAllUsers() ([]entity.Users, error) {
	var users []entity.Users
	err := userService.db.Transaction(func(tx *gorm.DB) error {
		var err error
		users, err = userService.userRepository.FindAll()
		if err != nil {
			return err
		}
		return nil
	})
	return users, err
}

func (userService *UserServiceImpl) FindOneByEmail(email string) (interface{}, error) {
	var user entity.Users
	err := userService.db.Transaction(func(tx *gorm.DB) error {
		u, err := userService.userRepository.FindByEmail(email)
		if err != nil {
			return err
		}
		user = *u.(*entity.Users)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, err
}

func (userService *UserServiceImpl) LoginUser(request dto.LoginRequest) error {
	err := userService.db.Transaction(func(tx *gorm.DB) error {
		u, err := userService.userRepository.FindByEmail(request.Email)
		if err != nil {
			return err
		}
		if u == nil {
			return errors.New("user not found")
		}
		user := *u.(*entity.Users)
		if user.Password != request.Password {
			return errors.New("password가 일치하지 않습니다")
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserServiceImpl) DeleteUserByEmail(email string) error {
	return userService.db.Transaction(func(tx *gorm.DB) error {
		err := userService.userRepository.DeleteByEmail(email)
		if err != nil {
			return err
		}
		return nil
	})
}
