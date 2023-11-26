package models

import (
	"github.com/jinzhu/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionId uint
	ProductId     *uint
	EventId       *uint
	UserId        *uint
	Qty           int32
	Status        string
	Product       Product `gorm:"foreignKey:ID;references:ProductId"`
	Users         Users   `gorm:"foreignKey:ID;references:UserId"`
	// Event even `gorm:"foreignKey:ID;references:ProductId"`
}
