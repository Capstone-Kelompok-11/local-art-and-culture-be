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
	}
}

func ConvertFromModelToArticleRes(data models.Article) *response.Article {
	return &response.Article{
		Id:      data.ID,
		Title:   data.Title,
		Content: data.Content,
		AdminId: data.AdminId,
	}
}

// func ConvertFromModelToAdminArticleRes(data models.SuperAdmin) *response.SuperAdminArticle {
//     return &response.SuperAdminArticle{
//         Id:          data.ID,
//         Name:        data.Name,
//         Email:       data.Email,
//         PhoneNumber: data.PhoneNumber,
//         Article:    *ConvertFromModelToArticleRes(data.Article), 
//     }
// }
