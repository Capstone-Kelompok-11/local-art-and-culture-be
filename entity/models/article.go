package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title     string
	AdminId   uint
	Content   string
	FilesId   *uint
	Status    string
	TotalLike uint
	Files     Files      `gorm:"foreignKey:FilesId"`
	Like      []Like     `gorm:"foreignKey:SourceId"`
	Admin     SuperAdmin `gorm:"foreignKey:AdminId"`
}
