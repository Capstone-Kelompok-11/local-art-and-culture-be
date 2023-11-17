package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

func ConvertFromShippingReqToModel(data request.Shipping) *models.SuperAdmin {
	return &models.SuperAdmin{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name:        data.Name,
	}
}

func ConvertFromModelToShippingRes(data models.Shipping) *response.Shipping {
	return &response.Shipping{
		Id:          data.ID,
		Name:        data.Name,
	}
}
