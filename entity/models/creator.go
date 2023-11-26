package models

import "github.com/jinzhu/gorm"

type Creator struct {
	gorm.Model
	OutletName  string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	UserId      uint   `gorm:"not null"`
	RoleId      uint   `gorm:"not null"`
	AddressId   *uint
	Users       Users    `gorm:"foreignKey:ID;references:UserId"`
	Role       Role     `gorm:"foreignKey:ID;references:RoleId"`
	Addresses   *Address `gorm:"foreignKey:ID;references:AddressId"`
}