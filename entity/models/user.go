package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Users struct {
	gorm.Model
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Username	string
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	PhoneNumber string    `gorm:"unique;not null"`
	NIK         string    `gorm:"not null"`
	Gender      string    `gorm:"not null"`
	BirthDate   time.Time `gorm:"not null"`
	RoleId      uint      `gorm:"not null"`
	Role        Role      `gorm:"foreignKey:RoleId"`
}