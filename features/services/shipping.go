package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"

	"lokasani/helpers/errors"
)

type IShippingService interface {
	CreateShippingMethod(data *request.Shipping) (response.Shipping, error)
	GetAllShipping() ([]response.Shipping, error)
	GetShipping(id string) (response.Shipping, error)
	UpdateShipping(id string, input request.Shipping) (response.Shipping, error)
	DeleteShipping(id string) (response.Shipping, error)
}

type ShippingService struct {
	shippingRepository repositories.IShippingRepository
}

func NewShippingService(repo repositories.IShippingRepository) *ShippingService {
	return &ShippingService{shippingRepository: repo}
}

func (sms *ShippingService) CreateShippingMethod(data *request.Shipping) (response.Shipping, error) {
	if data.Name == "" {
		return response.Shipping{}, errors.ERR_ROLE_IS_EMPTY
	}

	res, err := sms.shippingRepository.CreateShippingMethod(data)
	if err != nil {
		return response.Shipping{}, errors.ERR_CREATE_ROLE
	}

	return res, nil
}

func (sms *ShippingService) GetAllShipping() ([]response.Shipping, error) {
	err, res := sms.shippingRepository.GetAllShipping()
	if err != nil {
		return err, nil
	}
	return nil, res
}

func (sms *ShippingService) GetShipping(id string) (response.Shipping, error) {
	if id == "" {
		return response.Shipping{}, errors.ERR_GET_SHIPPING_BAD_REQUEST_ID
	}
	res, err := sms.shippingRepository.GetShipping(id)
	if err != nil {
		return response.Shipping{}, err
	}
	return res, nil
}

func (sms *ShippingService) UpdateShipping(id string, data request.Shipping) (response.Shipping, error) {
	if id == "" {
		return response.Shipping{}, errors.ERR_GET_SHIPPING_BAD_REQUEST_ID
	}
	res, err := sms.shippingRepository.UpdateShipping(id, data)
	if err != nil {
		return response.Shipping{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (sms *ShippingService) DeleteShipping(id string) (response.Shipping, error) {
	if id == "" {
		return response.Shipping{}, errors.ERR_GET_SHIPPING_BAD_REQUEST_ID
	}
	res, err := sms.shippingRepository.DeleteShipping(id)
	if err != nil {
		return response.Shipping{}, errors.ERR_DELETE_SHIPPING
	}
	return res, nil
}