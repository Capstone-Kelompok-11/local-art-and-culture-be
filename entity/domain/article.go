package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromArticleReqToModel(data request.Article) *models.Article {
	return &models.Article{
		Model: gorm.Model{
			ID: data.Id,
		},
		Title:   data.Title,
		Content: data.Content,
		AdminId: data.AdminId,
		Status:  data.Status,
		FilesId: data.FilesId,
	}
}

func ConvertFromModelToArticleRes(data models.Article) *response.Article {
	return &response.Article{
		Id:        data.ID,
		CreatedAt: data.CreatedAt,
		Title:     data.Title,
		Content:   data.Content,
		AdminId:   data.AdminId,
		TotalLike: data.TotalLike,
		Status:    data.Status,
		FilesId:   data.FilesId,
		Admin:     *ConvertFromModelToAdminRes(data.Admin),
		Like:      *ConvertFromModelsToLikeRes(data.Like),
	}
}
