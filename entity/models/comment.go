package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Comment 	string 	`gorm:"not null"`
	ArticleId	uint	`gorm:"not null"`
	LikeId		*uint	`gorm:"not null"`
}