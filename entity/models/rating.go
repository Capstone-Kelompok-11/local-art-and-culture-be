package models

import "github.com/jinzhu/gorm"

type Rating struct {
	gorm.Model
	Rating			uint16
	Ulasan			string
	UserId			uint
	ProductId		uint
	Users Users 	`gorm:"foreignKey:ID;references:UserId"`
	Product Product `gorm:"foreignKey:ID;references:ProductId"`
}