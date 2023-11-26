package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IEventRepository interface {
	CreateEvent(data *request.Event) (response.Event, error)
	GetAllEvent() ([]response.Event, error)
	GetEvent(id string) (response.Event, error)
	UpdateEvent(id string, input request.Event) (response.Event, error)
	DeleteEvent(id string) (response.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db}
}

func (er *eventRepository) CreateEvent(data *request.Event) (response.Event, error) {
	dataEvent := domain.ConvertFromEventReqToModel(*data)
	err := er.db.Create(&dataEvent).Error
	if err != nil {
		return response.Event{}, err
	}
	err = er.db.Preload("Category").Preload("Creator").First(&dataEvent, "id = ?", dataEvent.ID).Error
	return *domain.ConvertFromModelToEventRes(*dataEvent), nil
}

func (er *eventRepository) GetAllEvent() ([]response.Event, error) {
	var allEvent []models.Event
	var resAllEvent []response.Event
	err := er.db.Preload("Category").Preload("Creator").Find(&allEvent).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allEvent); i++ {
		eventVm := domain.ConvertFromModelToEventRes(allEvent[i])
		resAllEvent = append(resAllEvent, *eventVm)
	}
	return resAllEvent, nil
}

func (er *eventRepository) GetEvent(id string) (response.Event, error) {
	var eventData models.Event
	err := er.db.Preload("Category").Preload("Creator").First(&eventData, "id = ?", id).Error

	if err != nil {
		return response.Event{}, err
	}

	return *domain.ConvertFromModelToEventRes(eventData), nil
}

func (er *eventRepository) UpdateEvent(id string, input request.Event) (response.Event, error) {
	eventData := models.Event{}
	err := er.db.First(&eventData, "id = ?", id).Error

	if err != nil {
		return response.Event{}, errors.ERR_GET_EVENT_BAD_REQUEST_ID
	}

	if input.EventName != "" {
		eventData.EventName = input.EventName
	}
	if input.EventDescription != "" {
		eventData.EventDescription = input.EventDescription
	}
	if input.TicketPrice != 0 {
		eventData.TicketPrice = input.TicketPrice
	}
	if !input.FromDate.IsZero() {
		eventData.FromDate = input.FromDate
	}
	if !input.ToDate.IsZero() {
		eventData.ToDate = input.ToDate
	}
	if !input.StartTime.IsZero() {
		eventData.StartTime = input.StartTime
	}
	if !input.EndTime.IsZero() {
		eventData.EndTime = input.EndTime
	}

	if err = er.db.Save(&eventData).Error; err != nil {
		return response.Event{}, errors.ERR_UPDATE_DATA
	}
	return *domain.ConvertFromModelToEventRes(eventData), nil
}

func (er *eventRepository) DeleteEvent(id string) (response.Event, error) {
	eventData := models.Event{}
	res := response.Event{}
	find := er.db.Preload("Category").Preload("Creator").First(&eventData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToEventRes(eventData)
	}

	err := er.db.Delete(&eventData, "id = ?", id).Error
	if err != nil {
		return response.Event{}, err
	}
	return res, nil
}