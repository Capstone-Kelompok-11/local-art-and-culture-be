package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	Type      string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Description      string    `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	Qty       int       `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	EventId   uint      `gorm:"not null"`
	Event     Event     `gorm:"foreignKey:ID;references:EventId"`
}
