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
		PictureId: 	data.PictureId,
		EventId: 	data.EventId,
	}
}

func ConvertFromModelToGuestRes(data models.Guest) *response.Guest {
	return &response.Guest{
		Id: 		data.ID,
		Name: 		data.Name,
		Role: 		data.Role,
		PictureId: 	data.PictureId,
		EventId: 	data.EventId,
		Event: 		*ConvertFromModelToEventRes(data.Event),
	}
}