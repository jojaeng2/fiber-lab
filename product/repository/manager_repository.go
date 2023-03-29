package repository

import (
	"errors"
	"product/model"

	"gorm.io/gorm"
)

type ManagerRepository interface {
	Save(manager *model.Manager) error
	FindOneById(id int) (interface{}, error)
}

type ManagerRepositoryImpl struct {
	db *gorm.DB
}

func NewManagerRepository(db *gorm.DB) ManagerRepository {
	return &ManagerRepositoryImpl{db}
}

func (repository *ManagerRepositoryImpl) Save(manager *model.Manager) error {
	return repository.db.Create(manager).Error
}

func (repository *ManagerRepositoryImpl) FindOneById(id int) (interface{}, error) {
	var manager model.Manager
	err := repository.db.Where("id = ?", id).First(&manager).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("manager does not exist")
		}
		return nil, err
	}
	return manager, nil
}
