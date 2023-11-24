package models

import (
	"github.com/jinzhu/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionId uint
	ProductId     *uint
	EventId       *uint
	Qty           int32
	Status        string
	Product       Product `gorm:"foreignKey:ID;references:ProductId"`
	Event         Event   `gorm:"foreignKey:ID;references:EventId"`
}
