package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string
	AdminId uint
	Content string
	Like	[]Like	 `gorm:"foreignKey:SourceId"`
	Admin   SuperAdmin `gorm:"foreignKey:ID;references:AdminId"`
}