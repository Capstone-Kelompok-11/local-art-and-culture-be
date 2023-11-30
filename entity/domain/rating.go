package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromRatingReqToModel(data request.Rating) *models.Rating {
	return &models.Rating{
		Model: gorm.Model{
			ID: data.Id,
		},
		Rating: 	data.Rating,
		Ulasan: 	data.Ulasan,
		UserId: 	data.UserId,
		ProductId: 	data.ProductId,
	}
}

func ConvertFromModelToRatingRes(data models.Rating) *response.Rating {
	return &response.Rating{
		Id: data.ID,
		Rating: 	data.Rating,
		Ulasan: 	data.Ulasan,
		UserId: 	data.UserId,
		ProductId: 	data.ProductId,
		Users: 		*ConvertFromModelToUsersRes(data.Users),
		Product: 	*ConvertFromModelToProductsRes(data.Product),
	}
}