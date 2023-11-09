package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromUserReqToModel(data request.UserRequest) *models.Users {
	return &models.Users{
		Model: gorm.Model{
			ID: data.Id,
		},
	}
}