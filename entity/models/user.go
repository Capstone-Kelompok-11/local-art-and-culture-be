package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Users struct {
	gorm.Model
	FirstName		string
	LastName		string
	Email			string
	Password		string
	AlamatID		*uint
	NoHP			string
	BirthDate		time.Time
}