package models

import "github.com/jinzhu/gorm"

type Guest struct {
	gorm.Model
	Name 		string	`gorm:"not null"`
	Role		string	`gorm:"not null"`
	PictureId	*uint	`gorm:"not null"`
	EventId 	uint	`gorm:"not null"`
	Event	Event		`gorm:"foreignKey:ID;references:EventId"`
}