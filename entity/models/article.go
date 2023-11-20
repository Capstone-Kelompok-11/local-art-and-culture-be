package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string	`gorm:"not null"`
	AdminId uint	`gorm:"not null"`
	Content string	`gorm:"not null"`
}
