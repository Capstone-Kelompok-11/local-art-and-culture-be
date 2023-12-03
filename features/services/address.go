package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IAddressService interface {
	CreateAddress(data *request.Address) (response.Address, error)
	GetAllAddress(nameFilter string) ([]response.Address, error)
	GetAddress(id string) (response.Address, error)
	UpdateAddress(id string, data request.Address) (response.Address, error)
	DeleteAddress(id string) (response.Address, error)
}

type AddressService struct {
	AddressRepository repositories.IAddressRepository
}

func NewAddressService(repo repositories.IAddressRepository) *AddressService {
	return &AddressService{AddressRepository: repo}
}

func (ad *AddressService) CreateAddress(data *request.Address) (response.Address, error) {
	if data.Address == "" {
		return response.Address{}, errors.ERR_ADDRESS_IS_EMPTY
	}

	res, err := ad.AddressRepository.CreateAddress(data)
	if err != nil {
		return response.Address{}, errors.ERR_CREATE_ADDRESS_DATABASE
	}

	return res, nil
}

func (ad *AddressService) GetAllAddress(nameFilter string) ([]response.Address, error) {
	res, err := ad.AddressRepository.GetAllAddress(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (ad *AddressService) GetAddress(id string) (response.Address, error) {
	if id == "" {
		return response.Address{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := ad.AddressRepository.GetAddress(id)
	if err != nil {
		return response.Address{}, errors.ERR_GET_DATA
	}
	return res, nil
}

func (ad *AddressService) UpdateAddress(id string, data request.Address) (response.Address, error) {
	if id == "" {
		return response.Address{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := ad.AddressRepository.UpdateAddress(id, data)
	if err != nil {
		return response.Address{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (ad *AddressService) DeleteAddress(id string) (response.Address, error) {
	if id == "" {
		return response.Address{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := ad.AddressRepository.DeleteAddress(id)

	if err != nil {
		return response.Address{}, errors.ERR_DELETE_ADDRESS
	}

	return res, nil
}