package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	UserId    	uint `gorm:"not null"`
	SourceId  	uint `gorm:"not null"`
	SourceStr	string
	Active    	bool
}