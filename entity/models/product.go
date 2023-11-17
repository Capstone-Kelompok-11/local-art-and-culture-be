package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name 		string 	`gorm:"unique;not null"`
	Price		string 	`gorm:"unique;not null"`
	Description	string 	`gorm:"unique;not null"`
	Status		string 	`gorm:"unique;not null"`
	//AddressId	uint   	`gorm:"not null"`
	CategoryId	uint	`gorm:"not null"`
	CreatorId 	uint	`gorm:"not null"`
}