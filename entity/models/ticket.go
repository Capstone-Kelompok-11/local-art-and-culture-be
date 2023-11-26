package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	Name		string		`gorm:"not null"`
	Email		string		`gorm:"not null"`
	Type		string		`gorm:"not null"`
	Price		float64  	`gorm:"not null"`
	PhoneNumber string    	`gorm:"unique;not null"`
	Gender      string    	`gorm:"not null"`
	BirthDate   time.Time 	`gorm:"not null"`
	EventId		uint		`gorm:"not null"`
	Event 		Event		`gorm:"foreignKey:ID;references:EventId"`
}