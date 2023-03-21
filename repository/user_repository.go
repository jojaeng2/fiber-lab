package repository

import (
	"custom-modules/entity"
	"errors"
)

type UserRepository interface {
	Save(user entity.Users) error
	FindAll() ([]entity.Users, error)
	FindByEmail(email string) (interface{}, error)
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
	return nil
}

func (repo *UserRepositoryImpl) FindAll() ([]entity.Users, error) {
	return repo.users, nil
}

func (repo *UserRepositoryImpl) FindByEmail(email string) (interface{}, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
