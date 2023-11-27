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
		SourceId:  data.SourceId,
		SourceStr: data.SourceStr,
		UserId:    data.UserId,
	}
}

func ConvertFromModelToLikeRes(data models.Like) *response.Like {
	return &response.Like{
		Id:        data.ID,
		Active:    data.Active,
		SourceId:  data.SourceId,
		SourceStr: data.SourceStr,
		UserId:    data.UserId,
	}
}

func ConvertFromModelsToLikeRes(data []models.Like) *[]response.Like {
	var result []response.Like
	var temp response.Like
	for i := range data {
		temp.Id = data[i].ID
		temp.Active = data[i].Active
		temp.SourceId = data[i].SourceId
		temp.SourceStr = data[i].SourceStr
		result = append(result, temp)
	}
	return &result
}