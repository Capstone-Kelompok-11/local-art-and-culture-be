package models

import "github.com/jinzhu/gorm"

type Rating struct {
	gorm.Model
	Rating			uint16
	Ulasan			string
	UserId			uint
	ProductId		uint
	Users Users 	`gorm:"foreignKey:UserId"`
	Product Product `gorm:"foreignKey:ProductId"`
}