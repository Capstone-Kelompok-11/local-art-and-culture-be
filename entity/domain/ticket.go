package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromTicketReqToModel(data request.Ticket) *models.Ticket {
	return &models.Ticket{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name: 		 data.Name,
		Email: 		 data.Email,
		Type: 		 data.Type,
		Price: 		 data.Price,
		PhoneNumber: data.PhoneNumber,
		Gender: 	 data.Gender,
		BirthDate: 	 data.BirthDate,
		EventId: 	 data.EventId,
	}
}

func ConvertFromModelToTicketRes(data models.Ticket) *response.Ticket {
	return &response.Ticket{
		Id: 		 data.ID,
		Name: 		 data.Name,
		Email: 		 data.Email,
		Type: 		 data.Type,
		Price: 	 	 data.Price,
		PhoneNumber: data.PhoneNumber,
		Gender: 	 data.Gender,
		BirthDate: 	 data.BirthDate,
		EventId: 	 data.EventId,
		Event: 		 *ConvertFromModelToEventRes(data.Event),
	}
}