package service

import (
	"custom-modules/entity"
	"custom-modules/repository"
)

type UserService interface {
	SaveUser(user entity.Users) error
	FindAllUsers() error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (userService *UserServiceImpl) SaveUser(user entity.Users) error {
	return userService.userRepository.Save(user)
}

func (userService *UserServiceImpl) FindAllUsers() ([]entity.Users, error) {
	users, err := userService.userRepository.FindAll()
	if err != nil {
		return users, nil
	}
	return nil, err
}
