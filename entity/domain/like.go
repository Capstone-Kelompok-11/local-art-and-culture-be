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
		Like:      data.Like,
		ArticleId: data.ArticleId,
		LikeId:    data.LikeId,
	}
}

func ConvertFromModelToLikeRes(data models.Like) *response.Like {
	return &response.Like{
		ID:        data.Id,
		Like:      data.Like,
		ArticleId: data.ArticleId,
		LikeId:    data.LikeId,
	}
}
