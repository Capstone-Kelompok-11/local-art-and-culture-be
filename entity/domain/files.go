package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromFilesReqToModel(data request.Files, imageUrl string) *models.Files {
	return &models.Files{
		Model: gorm.Model{
			ID: data.Id,
		},
		SourceId:  data.SourceId,
		SourceStr: data.SourceStr,
		Image: 	   imageUrl,
	}
}

func ConvertFromModelToFilesRes(data models.Files) *response.Files {
	return &response.Files{
		Id:		   data.ID,
		SourceId:  data.SourceId,
		SourceStr: data.SourceStr,
		Image: 	   data.Image,
	}
}
