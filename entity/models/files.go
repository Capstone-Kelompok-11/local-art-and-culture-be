package models

import "github.com/jinzhu/gorm"

type Files struct {
	gorm.Model
	SourceId  uint
	SourceStr string
	Url       string
	Filename  string
}
