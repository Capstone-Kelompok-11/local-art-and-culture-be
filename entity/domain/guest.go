package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromGuestReqToModel(data request.Guest) *models.Guest {
	return &models.Guest{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name: 		data.Name,
		Role: 		data.Role,
		//PictureId: 	data.PictureId,
		EventId: 	data.EventId,
	}
}

func ConvertFromModelToGuestRes(data models.Guest) *response.Guest {
	return &response.Guest{
		Id: 		data.ID,
		Name: 		data.Name,
		Role: 		data.Role,
		//PictureId: 	data.PictureId,
		EventId: 	data.EventId,
	}
}

func ConvertFromGuestReqToModels(data []request.Guest) *[]models.Guest {
	var result []models.Guest
	var temp models.Guest
	for i := range data {
		temp.ID = data[i].Id
		temp.Name = data[i].Name
		temp.Role = data[i].Role
		temp.EventId = data[i].EventId
		result = append(result, temp)
	}
	return &result
}

func ConvertFromModelsToGuestRes(data []models.Guest) *[]response.Guest {
	var result []response.Guest
	var temp response.Guest
	for i := range data {
		temp.Id = data[i].ID
		temp.Name = data[i].Name
		temp.Role = data[i].Role
		temp.EventId = data[i].EventId
		result = append(result, temp)
	}
	return &result
}