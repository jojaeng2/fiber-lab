package controller

import (
	"custom-modules/dto"
	"custom-modules/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BoardController interface {
	SaveBoard(c *fiber.Ctx) error
	FindOneByTitle(c *fiber.Ctx) error
	FindOneById(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
}

type BoardControllerImpl struct {
	boardService service.BoardService
}

func NewBoardController(boardService service.BoardService) BoardController {
	return &BoardControllerImpl{
		boardService,
	}
}

func (controller *BoardControllerImpl) SaveBoard(c *fiber.Ctx) error {
	var request dto.BoardCreateDto
	err := c.BodyParser(&request)

	if err != nil {
		return err
	}

	controller.boardService.Save(request)
	return c.SendStatus(fiber.StatusOK)
}

func (controller *BoardControllerImpl) FindOneByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	board, err := controller.boardService.FindByTitle(title)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": title + "과 일치하는 게시글이 존재하지 않습니다.",
		})
	}
	return c.JSON(board)
}

func (controller *BoardControllerImpl) FindOneById(c *fiber.Ctx) error {
	id, err2 := strconv.Atoi(c.Params("id"))
	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id" + "는 숫자만 포함되어야 합니다.",
		})
	}
	board, err := controller.boardService.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "id와 일치하는 게시글이 존재하지 않습니다.",
		})
	}
	return c.JSON(board)
}

func (controller *BoardControllerImpl) DeleteById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id" + "는 숫자만 포함되어야 합니다.",
		})
	}
	err2 := controller.boardService.Delete(id)
	if err2 != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "존재하지 않는 게시글을 삭제할 수 없습니다.",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
