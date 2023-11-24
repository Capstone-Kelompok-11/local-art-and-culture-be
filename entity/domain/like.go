package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromLikeReqToModel(data request.Like) *models.Like {
	return &models.Like{
		Model: gorm.Model{
			ID: data.Id,
		},
		Active:    data.Active,
		ArticleId: data.ArticleId,
		UserId:    data.UserId,
	}
}

func ConvertFromModelToLikeRes(data models.Like) *response.Like {
	return &response.Like{
		Id:        data.ID,
		Active:    data.Active,
		ArticleId: data.ArticleId,
		UserId:    data.UserId,
	}
}
