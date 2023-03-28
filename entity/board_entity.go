package entity

import (
	"custom-modules/entity"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Title       string `gorm:"size:255"`
	Description string
	Comments    []entity.Comment
}
