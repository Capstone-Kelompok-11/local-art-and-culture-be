package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ArticleId uint `gorm:"not null"`
	Active    bool
}