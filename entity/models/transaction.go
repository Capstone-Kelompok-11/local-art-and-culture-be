package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionDate  time.Time
	UserId           uint
	PaymentMethodId  uint
	ShippingMethodId uint
	Status           string
	User             Users    `gorm:"foreignKey:ID;references:UserId"`
	Payment          Payment  `gorm:"foreignKey:ID;references:PaymentMethodId"`
	Shipping         Shipping `gorm:"foreignKey:ID;references:ShippingMethodId"`
}