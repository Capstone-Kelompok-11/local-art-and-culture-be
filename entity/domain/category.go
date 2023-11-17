package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromCategoryReqToModel(data request.Category) *models.Category {
	return &models.Category{
		Model: gorm.Model{
			ID: data.Id,
		},
		Category: 	data.Category,
		Type: 		data.Type,
	}
}

func ConvertFromModelToCategoryRes(data models.Category) *response.Category {
	return &response.Category{
		Id: 		data.ID,
		Category: 	data.Category,
		Type: 		data.Type,
	}
}