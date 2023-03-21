package service

import (
	"custom-modules/dto"
	"custom-modules/entity"
	"custom-modules/repository"
)

type UserService interface {
	SaveUser(request dto.CreateUserRequest) error
	FindAllUsers() ([]entity.Users, error)
	FindOneByEmail(email string) (interface{}, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (userService *UserServiceImpl) SaveUser(request dto.CreateUserRequest) error {
	user := entity.Users{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	return userService.userRepository.Save(user)
}

func (userService *UserServiceImpl) FindAllUsers() ([]entity.Users, error) {
	users, err := userService.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService *UserServiceImpl) FindOneByEmail(email string) (interface{}, error) {
	user, err := userService.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}
