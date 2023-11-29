package models

import "github.com/jinzhu/gorm"

type Wishlist struct {
	gorm.Model
	Quantity  uint
	ProductId uint
	TicketId  uint
	Product   Product `gorm:"foreignKey:ProductId"`
	Ticket    Ticket  `gorm:"foreignKey:TicketId"`
}
