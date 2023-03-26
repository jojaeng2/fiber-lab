package repository

import (
	"custom-modules/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.Users) error
	FindAll() ([]entity.Users, error)
	FindByEmail(email string) (interface{}, error)
	DeleteByEmail(email string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repo *UserRepositoryImpl) Save(user *entity.Users) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepositoryImpl) FindAll() ([]entity.Users, error) {
	var users []entity.Users
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepositoryImpl) FindByEmail(email string) (interface{}, error) {
	var user entity.Users
	err := repo.db.Where("email = ?", email).First(&user).Error
	fmt.Println(email, err)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) DeleteByEmail(email string) error {
	var user entity.Users
	err := repo.db.Where("email = ?", email).First(&user).Error
	fmt.Println(email, err)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return err
	}
	repo.db.Delete(&user)
	return nil
}
