package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	FromDate         time.Time `gorm:"not null"`
	ToDate           time.Time `gorm:"not null"`
	EventName        string    `gorm:"not null"`
	StartTime        time.Time `gorm:"not null"`
	EndTime          time.Time `gorm:"not null"`
	EventDescription string    `gorm:"not null"`
	TicketPrice      float64   `gorm:"not null"`
	AddressId        *uint
	CreatorId        uint     `gorm:"not null"`
	CategoryId       uint     `gorm:"not null"`
	Creator          Creator  `gorm:"foreignKey:ID;references:CreatorId"`
	Category         Category 	`gorm:"foreignKey:ID;references:CategoryId"`
	Guest			[]Guest		`gorm:"foreignKey:EventId"`
}
