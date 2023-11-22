package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromPaymentReqToModel(data request.Payment) *models.Payment {
	return &models.Payment{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name: data.Name,
	}
}

func ConvertFromModelToPaymentRes(data models.Payment) *response.Payment {
	return &response.Payment{
		Id: 	data.ID,
		Name: 	data.Name,
	}
}