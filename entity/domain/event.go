package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromEventReqToModel(data request.Event) *models.Event {
	return &models.Event{
		Model: gorm.Model{
			ID: data.Id,
		},
		EventName:        data.EventName,
		EventDescription: data.EventDescription,
		FromDate:         data.FromDate,
		ToDate:           data.ToDate,
		AddressId:        data.AddressId,
		CategoryId:       data.CategoryId,
		CreatorId:        data.CreatorId,
		Creator:          *ConvertFromCreatorReqToModel(data.Creator),
		Category:         *ConvertFromCategoryReqToModel(data.Category),
	}
}

func ConvertFromModelToEventRes(data models.Event) *response.Event {
	return &response.Event{
		Id:               data.ID,
		EventName:        data.EventName,
		EventDescription: data.EventDescription,
		FromDate:         data.FromDate,
		ToDate:           data.ToDate,
		AddressId:        data.AddressId,
		CategoryId:       data.CategoryId,
		CreatorId:        data.CreatorId,
		Creator:          *ConvertFromModelToCreatorRes(data.Creator),
		Category:         *ConvertFromModelToCategoryRes(data.Category),
		Guest: 			  *ConvertFromModelsToGuestRes(data.Guest),
	}
}

func ConvertFromModelToEventsRes(data models.Event) *response.Events {
	return &response.Events{
		Id:               data.ID,
		EventName:        data.EventName,
		EventDescription: data.EventDescription,
		FromDate:         data.FromDate,
		ToDate:           data.ToDate,
		AddressId:        data.AddressId,
		CategoryId:       data.CategoryId,
		CreatorId:        data.CreatorId,
	}
}