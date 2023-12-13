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
		Stock: 		data.Stock,
		Price:       data.Price,
		Description: data.Description,
		Status:      data.Status,
		TotalProduct: data.TotalProduct,
		CategoryId:  data.CategoryId,
		CreatorId:   data.CreatorId,
	}
}

func ConvertFromModelToProductRes(data models.Product) *response.Product {
	return &response.Product{
		Id:          data.ID,
		Name:        data.Name,
		Stock: 		data.Stock,
		Price:       data.Price,
		Description: data.Description,
		Status:      data.Status,
		TotalProduct: data.TotalProduct,
		CategoryId:  data.CategoryId,
		CreatorId:   data.CreatorId,
		Category:    *ConvertFromModelToCategoryRes(data.Category),
		Creator:     *ConvertFromModelToCreatorRes(data.Creator),
	}
}

func ConvertFromModelToProductsRes(data models.Product) *response.Products {
	return &response.Products{
		Id:          data.ID,
		Name:        data.Name,
		Stock: 		data.Stock,
		Price:       data.Price,
		Description: data.Description,
		Status:      data.Status,
		TotalProduct: data.TotalProduct,
		CategoryId:  data.CategoryId,
		CreatorId:   data.CreatorId,
		TotalLike: 	 data.TotalLike,
		Category:    *ConvertFromModelToCategoryRes(data.Category),
		Like: 	 *ConvertFromModelsToLikeRes(data.Like),
	}
}