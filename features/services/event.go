package services

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
	"math"
)

type IEventService interface {
	CreateEvent(data *request.Event) (response.Event, error)
	GetAllEvent(nameFilter, startDate, endDate string, page, pageSize int) ([]response.Event, int, error)
	GetAllAvailableEvent(nameFilter, startDate, endDate string, page, pageSize int) ([]response.Event, int, error)
	GetEvent(id string) (response.Event, error)
	UpdateEvent(id string, input request.Event) (response.Event, error)
	DeleteEvent(id string) (response.Event, error)
	CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	GetNextPage(currentPage, allPages int) int
	GetPrevPage(currentPage int) int
}

type EventService struct {
	eventRepository repositories.IEventRepository
}

func NewEventService(repo repositories.IEventRepository) *EventService {
	return &EventService{eventRepository: repo}
}

func (er *EventService) CreateEvent(data *request.Event) (response.Event, error) {
	if data.EventName == "" {
		fmt.Println(data)
		return response.Event{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.EventDescription == "" {
		fmt.Println(data)
		return response.Event{}, errors.ERR_DESCRIPTION_IS_EMPTY
	}
	if data.FromDate.IsZero() || data.ToDate.IsZero() {
		fmt.Println(data)
		return response.Event{}, errors.ERR_EVENT_DATE_IS_EMPTY
	}

	res, err := er.eventRepository.CreateEvent(data)
	if err != nil {
		return response.Event{}, errors.ERR_CREATE_EVENT_DATABASE
	}
	fmt.Println(data)
	return res, nil
}

func (er *EventService) GetAllEvent(nameFilter, startDate, endDate string, page, pageSize int) ([]response.Event, int, error) {
	res, allItems, err := er.eventRepository.GetAllEvent(nameFilter, startDate, endDate, page, pageSize)
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}
	return res, allItems, nil
}

func (er *EventService) GetAllAvailableEvent(nameFilter, startDate, endDate string, page, pageSize int) ([]response.Event, int, error) {
	res, allItems, err := er.eventRepository.GetAllAvailableEvent(nameFilter, startDate, endDate, page, pageSize)
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}
	return res, allItems, nil
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

func (er *EventService) CalculatePaginationValues(page, pageSize, allItmes int) (int, int) {
	pageInt := page
	if pageInt <= 0 {
		pageInt = 1
	}

	allPages := int(math.Ceil(float64(allItmes) / float64(pageSize)))

	if pageInt > allPages {
		pageInt = allPages
	}

	return pageInt, allPages
}

func (er *EventService) GetNextPage(currentPage, allPages int) int {
	if currentPage < allPages {
		return currentPage + 1
	}
	return allPages
}

func (er *EventService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}