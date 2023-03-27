package repository

import (
	"custom-modules/board/entity"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Save(board *entity.Board) error
	FindByTitle(title string) (interface{}, error)
	FindById(id int) (interface{}, error)
	DeleteById(id int) error
}

type BoardRepositoryImpl struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &BoardRepositoryImpl{db}
}

func (repo *BoardRepositoryImpl) Save(board *entity.Board) error {
	return repo.db.Create(board).Error
}

func (repo *BoardRepositoryImpl) FindByTitle(title string) (interface{}, error) {
	var board entity.Board
	err := repo.db.Where("title = ?", title).First(&board).Error
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (repo *BoardRepositoryImpl) FindById(id int) (interface{}, error) {
	var board entity.Board
	err := repo.db.Where("id = ?", id).First(&board).Error
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (repo *BoardRepositoryImpl) DeleteById(id int) error {
	var board entity.Board
	err := repo.db.Where("id = ?", id).First(&board).Error
	if err != nil {
		return err
	}
	repo.db.Delete(&board)
	return nil
}
