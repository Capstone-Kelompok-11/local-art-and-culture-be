package models

import "gorm.io/gorm"

type SuperAdmin struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	PhoneNumber string `gorm:"unique;not null"`
	RoleId      uint   `gorm:"not null"`
	Role        Role   `gorm:"foreignKey:RoleId"`
}