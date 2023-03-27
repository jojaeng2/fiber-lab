package service

import (
	"custom-modules/board/dto"
	"custom-modules/board/entity"
	"custom-modules/board/repository"

	"gorm.io/gorm"
)

type BoardService interface {
	Save(request dto.BoardCreateDto) error
	FindByTitle(title string) (interface{}, error)
	FindById(id int) (interface{}, error)
	Delete(id int) error
}

type BoardServiceImpl struct {
	boardRepository repository.BoardRepository
	db              *gorm.DB
}

func NewBoardService(boardRepository repository.BoardRepository, db *gorm.DB) BoardService {
	return &BoardServiceImpl{
		boardRepository: boardRepository,
		db:              db,
	}
}

func (boardService *BoardServiceImpl) Save(request dto.BoardCreateDto) error {
	return boardService.db.Transaction(func(tx *gorm.DB) error {
		board := entity.Board{
			Title:       request.Title,
			Description: request.Description,
		}
		return boardService.boardRepository.Save(&board)
	})
}

func (boardService *BoardServiceImpl) FindByTitle(title string) (interface{}, error) {
	var board entity.Board
	err := boardService.db.Transaction(func(tx *gorm.DB) error {
		b, err := boardService.boardRepository.FindByTitle(title)
		if err != nil {
			return err
		}
		board = *b.(*entity.Board)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (boardService *BoardServiceImpl) FindById(id int) (interface{}, error) {
	var board entity.Board
	err := boardService.db.Transaction(func(tx *gorm.DB) error {
		b, err := boardService.boardRepository.FindById(id)
		if err != nil {
			return err
		}
		board = *b.(*entity.Board)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (boardService *BoardServiceImpl) Delete(id int) error {
	err := boardService.db.Transaction(func(tx *gorm.DB) error {
		b, err := boardService.boardRepository.FindById(id)
		if err != nil {
			return err
		}
		boardService.boardRepository.Delete(b.(*entity.Board))
		return nil
	})

	return err
}
