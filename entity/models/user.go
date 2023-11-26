package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Users struct {
	gorm.Model
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	PhoneNumber string    `gorm:"unique;not null"`
	BirthDate   time.Time `gorm:"not null"`
}
