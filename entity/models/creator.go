package models

import "github.com/jinzhu/gorm"

type Creator struct {
	gorm.Model
	OutletName  string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	UserId      uint   `gorm:"not null"`
	RoleId 		uint   `gorm:"not null"`
	AddressId   *uint
	Address     string
	Users       Users    `gorm:"foreignKey:UserId"`
	Role        Role     `gorm:"foreignKey:RoleId"`
	Addresses   *Address `gorm:"foreignKey:AddressId"`
}
