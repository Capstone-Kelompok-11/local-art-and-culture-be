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
		Type:        data.Type,
		Price:       data.Price,
		EventId:     data.EventId,
		Qty:         data.Qty,
		Name:        data.Name,
		Description: data.Description,
	}
}

func ConvertFromModelToTicketRes(data models.Ticket) *response.Ticket {
	return &response.Ticket{
		Id:          data.ID,
		Type:        data.Type,
		Price:       data.Price,
		Qty:         data.Qty,
		EventId:     data.EventId,
		Name:        data.Name,
		Description: data.Description,
		Event:       *ConvertFromModelToEventsRes(data.Event),
	}
}
