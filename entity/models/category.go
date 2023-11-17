package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Category string `gorm:"not null"`
	Type     string `gorm:"not null"`
}