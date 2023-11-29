package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromFilesReqToModel(data request.Files) *models.Files {
	return &models.Files{
		Model: gorm.Model{
			ID: data.Id,
		},
		SourceId:  data.SourceId,
		SourceStr: data.SourceStr,
		Filename:  data.Filename,
		Url:       data.Url,
	}
}

func ConvertFromModelToFilesRes(data models.Files) *response.Files {
	return &response.Files{
		SourceId:  data.SourceId,
		SourceStr: data.SourceStr,
		Filename:  data.Filename,
		Url:       data.Url,
	}
}
