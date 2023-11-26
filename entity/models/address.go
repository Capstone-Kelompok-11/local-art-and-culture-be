package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Address string
	Users   []Users   `gorm:"many2many:user_addresses;"`
}
