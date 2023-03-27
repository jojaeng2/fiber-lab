package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Description string
	BoardID     uint // 추가된 필드
}
