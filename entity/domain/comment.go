package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromCommentReqToModel(data request.Comment) *models.Comment{
	return &models.Comment{
		Model: gorm.Model{
			ID: data.Id,
		},
		Comment: 	data.Comment,
		ArticleId: 	data.ArticleId,
		LikeId: 	data.LikeId,
		UserId:		data.UserId,
	}
}

func ConvertFromModelToCommentRes(data models.Comment) *response.Comment{
	return &response.Comment{
		Id: data.ID,
		Comment: 	data.Comment,
		ArticleId: 	data.ArticleId,
		LikeId: 	data.LikeId,
		UserId: 	data.UserId,
		User: *ConvertFromModelToUserRes(data.User),
	}
}