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
	FindOneByEmail(email string) interface{}
	LoginUser(request dto.LoginRequest)
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

func (userService *UserServiceImpl) FindOneByEmail(email string) interface{} {
	return userService.userRepository.FindByEmail(email).OrElseThrow()
}

func (userService *UserServiceImpl) LoginUser(request dto.LoginRequest) {
	user := userService.userRepository.FindByEmail(request.Email)
	fmt.Println(user)

	u, ok := user.OrElseThrow().(*entity.Users) // 타입 어서션
	if !ok {
		panic(errors.New("unexpected user type"))
	}
	if u.Password != request.Password {
		panic(errors.New("잘못된 정보입니다"))
	}
}
