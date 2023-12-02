package models

import "github.com/jinzhu/gorm"

type SaveChatbot struct {
	gorm.Model
	Message  string
	Response string
	UserId   uint
}