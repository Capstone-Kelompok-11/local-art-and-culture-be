package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `gorm:"not null"`
	Stock 		bool 
	Price       float64  `gorm:"not null"`
	Description string   `gorm:"not null"`
	Status      string   `gorm:"not null"`
	TotalProduct int64
	CategoryId  uint     `gorm:"not null"`
	CreatorId   uint     `gorm:"not null"`
	Like        []Like   `gorm:"foreignKey:SourceId"`
	Creator     Creator  `gorm:"foreignKey:CreatorId"`
	Category    Category `gorm:"foreignKey:CategoryId"`
}
