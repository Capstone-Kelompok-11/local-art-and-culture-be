package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	Name 	string `gorm:"not null"`
}