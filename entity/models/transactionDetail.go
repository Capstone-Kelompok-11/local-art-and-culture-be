package models

import (
	"github.com/jinzhu/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionId uint
	ProductId     *uint
	TicketId      *uint
	Fullname      *string
	Contact       *string
	Qty           int32
	Product       Product `gorm:"foreignKey:ID;references:ProductId"`
	Ticket        Ticket  `gorm:"foreignKey:ID;references:TicketId"`
}
