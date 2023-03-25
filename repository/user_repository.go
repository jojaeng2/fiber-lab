package repository

import (
	"custom-modules/entity"
	"custom-modules/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.Users) error
	FindAll() ([]entity.Users, error)
	FindByEmail(email string) utils.Optional
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

func (repo *UserRepositoryImpl) FindByEmail(email string) utils.Optional {
	var user entity.Users
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return *utils.OfNullable(nil)
		}
		return *utils.OfNullable(nil)
	}
	return *utils.OfNullable(user)
}
