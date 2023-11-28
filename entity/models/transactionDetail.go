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
	Nik           *string
	Qty           int32
	Product       Product `gorm:"foreignKey:ProductId"`
	Ticket        Ticket  `gorm:"foreignKey:TicketId"`
}
