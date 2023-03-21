package repository

import (
	"custom-modules/entity"
	"fmt"
)

type UserRepository interface {
	Save(user entity.Users) error
	FindAll() ([]entity.Users, error)
}

type UserRepositoryImpl struct {
	users []entity.Users
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		users: []entity.Users{},
	}
}

func (repo *UserRepositoryImpl) Save(user entity.Users) error {
	repo.users = append(repo.users, user)
	fmt.Println(repo.users)
	return nil
}

func (repo *UserRepositoryImpl) FindAll() ([]entity.Users, error) {
	fmt.Println("repository = ", repo.users)

	return repo.users, nil
}
