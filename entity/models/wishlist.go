package models

import "github.com/jinzhu/gorm"

type Wishlist struct {
	gorm.Model
	Quantity 	uint
	ProductId 	uint
	TicketId 	uint
	Product Product	`gorm:"foreignKey:ID;references:ProductId"`
	Ticket 	Ticket	`gorm:"foreignKey:ID;references:TicketId"`
}