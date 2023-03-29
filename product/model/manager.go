package model

import (
	"gorm.io/gorm"
)

type Manager struct {
	gorm.Model
	Email    string
	Password string
	Name     string
}
