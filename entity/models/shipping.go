package models

import "gorm.io/gorm"

type Shipping struct {
	gorm.Model
	Name        string    `gorm:"not null"`
}