package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionDate   time.Time
	UserId            uint
	PaymentMethodId   uint
	ShippingMethodId  uint
	Status            string
	User              Users               `gorm:"foreignKey:UserId"`
	Payment           Payment             `gorm:"foreignKey:PaymentMethodId"`
	Shipping          Shipping            `gorm:"foreignKey:ShippingMethodId"`
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionId"`
}
