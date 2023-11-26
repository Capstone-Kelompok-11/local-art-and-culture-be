package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string   `gorm:"not null"`
	Price       float64  `gorm:"not null"`
	Description string   `gorm:"not null"`
	Status      string   `gorm:"not null"`
	CategoryId  uint     `gorm:"not null"`
	CreatorId   uint     `gorm:"not null"`
	Creator     Creator  `gorm:"foreignKey:ID;references:CreatorId"`
	Category    Category `gorm:"foreignKey:ID;references:CategoryId"`
}