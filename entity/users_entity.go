package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string
	Password string
}
