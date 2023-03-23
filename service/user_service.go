package service

import (
	"custom-modules/dto"
	"custom-modules/entity"
	"custom-modules/repository"
	"errors"
	"fmt"
)

type UserService interface {
	SaveUser(request dto.CreateUserRequest) error
	FindAllUsers() ([]entity.Users, error)
	FindOneByEmail(email string) (interface{}, error)
	LoginUser(request dto.LoginRequest) error
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
	return userService.userRepository.Save(&user)
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

func (userService *UserServiceImpl) LoginUser(request dto.LoginRequest) error {
	user, err := userService.userRepository.FindByEmail(request.Email)
	fmt.Println(user)

	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	u, ok := user.(*entity.Users) // 타입 어서션
	if !ok {
		return errors.New("unexpected user type")
	}
	if u.Password == request.Password {
		return nil
	}
	return errors.New("잘못된 정보입니다.")
}
