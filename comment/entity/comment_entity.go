package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Description string
	Board       uint
}
