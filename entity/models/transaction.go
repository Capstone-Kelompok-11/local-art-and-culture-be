package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionDate   time.Time
	TransactionNumber string
	UserId            uint
	PaymentMethodId   uint
	ShippingMethodId  *uint
	Status            string
	User              Users               `gorm:"foreignKey:UserId"`
	Payment           Payment             `gorm:"foreignKey:PaymentMethodId"`
	Shipping          Shipping            `gorm:"foreignKey:ShippingMethodId"`
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionId"`
}

type TransactionReport struct {
	Id              uint
	TransactionDate time.Time
	Status          string
	Qty             int32
	Price           float64
	Nominal         float64
}
