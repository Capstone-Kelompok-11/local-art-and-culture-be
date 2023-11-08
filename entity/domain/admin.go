package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

func ConvertFromAdminReqToModel(data request.Admin) *models.Admin {
	return &models.Admin{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	}
}

func ConvertFromModelToAdminRes(data models.Admin) *response.Admin{
	return &response.Admin{
		Id: data.ID,
		Name: data.Name,
		Email: data.Email,
		PhoneNumber: data.PhoneNumber,
	}
}
