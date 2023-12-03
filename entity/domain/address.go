package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromAddressReqToModel(data request.Address) *models.Address {
	return &models.Address{
		Model: gorm.Model{
			ID: data.Id,
		},
		Address:  data.Address,
		Location: data.Location,
	}
}

func ConvertFromModelToAddressRes(data models.Address) *response.Address {
	return &response.Address{
		Id: data.ID,
		Address: data.Address,
		Location: data.Location,
	}
}