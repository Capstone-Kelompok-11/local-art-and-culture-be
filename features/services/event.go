package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IEventService interface {
	CreateEvent(data *request.Event) (response.Event, error)
	GetAllEvent() ([]response.Event, error)
	GetEvent(id string) (response.Event, error)
	UpdateEvent(id string, input request.Event) (response.Event, error)
	DeleteEvent(id string) (response.Event, error)
}

type EventService struct {
	eventRepository repositories.IEventRepository
}

func NewEventService(repo repositories.IEventRepository) *EventService {
	return &EventService{eventRepository: repo}
}

func (er *EventService) CreateEvent(data *request.Event) (response.Event, error) {
	if data.EventName == "" {
		return response.Event{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.EventDescription == "" {
		return response.Event{}, errors.ERR_DESCRIPTION_IS_EMPTY
	}
	if data.FromDate.IsZero() || data.ToDate.IsZero() {
		return response.Event{}, errors.ERR_EVENT_DATE_IS_EMPTY
	}

	res, err := er.eventRepository.CreateEvent(data)
	if err != nil {
		return response.Event{}, errors.ERR_CREATE_EVENT_DATABASE
	}
	return res, nil
}

func (er *EventService) GetAllEvent() ([]response.Event, error) {
	res, err := er.eventRepository.GetAllEvent()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (er *EventService) GetEvent(id string) (response.Event, error) {
	if id == "" {
		return response.Event{}, errors.ERR_GET_EVENT_BAD_REQUEST_ID
	}
	res, err := er.eventRepository.GetEvent(id)
	if err != nil {
		return response.Event{}, err
	}
	return res, nil
}

func (er *EventService) UpdateEvent(id string, input request.Event) (response.Event, error) {
	if id == "" {
		return response.Event{}, errors.ERR_GET_EVENT_BAD_REQUEST_ID
	}
	res, err := er.eventRepository.UpdateEvent(id, input)
	if err != nil {
		return response.Event{}, err
	}
	return res, nil
}

func (er *EventService) DeleteEvent(id string) (response.Event, error) {
	if id == "" {
		return response.Event{}, errors.ERR_GET_EVENT_BAD_REQUEST_ID
	}
	res, err := er.eventRepository.DeleteEvent(id)
	if err != nil {
		return response.Event{}, err
	}
	return res, nil
}
