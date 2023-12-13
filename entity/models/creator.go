package models

import "github.com/jinzhu/gorm"

type Creator struct {
	gorm.Model
	OutletName  string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	UserId      uint   `gorm:"not null"`
	Roles      string   `gorm:"type:enum('event_creator', 'product_creator')"`
	AddressId   *uint
	Users       Users    `gorm:"foreignKey:UserId"`
	Role       Role     `gorm:"foreignKey:Roles"`
	Addresses   *Address `gorm:"foreignKey:AddressId"`
}