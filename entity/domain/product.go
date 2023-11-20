package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromProductReqToModel(data request.Product) *models.Product {
	return &models.Product{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
		Status:      data.Status,
		CategoryId:  data.CategoryId,
		CreatorId:   data.CreatorId,
	}
}

func ConvertFromModelToProductRes(data models.Product) *response.Product {
	return &response.Product{
		Id:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
		Status:      data.Status,
		CategoryId:  data.CategoryId,
		CreatorId:   data.CreatorId,
		Category:    *ConvertFromModelToCategoryRes(data.Category),
		Creator:     *ConvertFromModelToCreatorRes(data.Creator),
	}
}

// func ConvertFromModelToCategoryProductRes(data models.Creator) *response.Category{
// 	return &response.Category{
// 		Id: 	  	data.ID,
// 		Category: 	data.Category,
// 		Type: 	  	data.Type,
// 		Products:	*ConvertFromModelToProductRes(data.Products)
// 	}
// }
