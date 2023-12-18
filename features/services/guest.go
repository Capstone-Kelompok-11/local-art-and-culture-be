package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IGuestService interface {
	CreateGuest(data *request.Guest) (response.Guest, error)
	GetAllGuest(nameFilter string) ([]response.Event, error)
	GetGuest(id string) (response.Event, error)
	UpdateGuest(id string, input request.Guest) (response.Guest, error)
	DeleteGuest(id string) (response.Guest, error)
}

type GuestService struct {
	guestRepository repositories.IGuestRepository
}

func NewGuestService(repo repositories.IGuestRepository) *GuestService {
	return &GuestService{guestRepository: repo}
}

func (gu *GuestService) CreateGuest(data *request.Guest) (response.Guest, error) {
	if data.Name == "" {
		return response.Guest{}, errors.ERR_NAME_IS_EMPTY
	} 
	if data.Role == "" {
		return response.Guest{}, errors.ERR_ROLE_IS_EMPTY
	}

	res, err := gu.guestRepository.CreateGuest(data)
	if err != nil {
		return response.Guest{}, errors.ERR_CREATE_GUEST_DATABASE
	}
	return res, nil
}

func (gu *GuestService) GetAllGuest(nameFilter string) ([]response.Event, error) {
	res, err := gu.guestRepository.GetAllGuest(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (gu *GuestService) GetGuest(id string) (response.Event, error) {
	if id == "" {
		return response.Event{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := gu.guestRepository.GetGuest(id)
	if err != nil {
		return response.Event{}, err
	}
	return res, nil
}

func (gu *GuestService) UpdateGuest(id string, input request.Guest) (response.Guest, error) {
	if id == "" {
		return response.Guest{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := gu.guestRepository.UpdateGuest(id, input)
	if err != nil {
		return response.Guest{}, err
	}
	return res, nil
}

func (gu *GuestService) DeleteGuest(id string) (response.Guest, error) {
	if id == "" {
		return response.Guest{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := gu.guestRepository.DeleteGuest(id)
	if err != nil {
		return response.Guest{}, err
	}
	return res, nil
}