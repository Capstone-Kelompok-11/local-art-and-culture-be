package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"

	// "lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"
	// "lokasani/helpers/middleware"
)

type IShippingService interface {
	GetAllShipping() (error, []response.Shipping)
	GetShipping(id string) (error, response.Shipping)
	UpdateShipping(id string, input request.Shipping) (error, response.Shipping)
	DeleteShipping(id string) (error, response.Shipping)
}

type ShippingService struct {
	shippingRepository repositories.IShippingRepository
}

func NewShippingService(repo repositories.IShippingRepository) *ShippingService {
	return &ShippingService{shippingRepository: repo}
}

func (as *ShippingService) GetAllShipping() (error, []response.Shipping) {
	err, res := as.shippingRepository.GetAllShipping()
	if err != nil {
		return err, nil
	}
	return nil, res
}

func (as *ShippingService) GetShipping(id string) (error, response.Shipping) {
	if id == "" {
		return errors.ERR_GET_SHIPPING_BAD_REQUEST_ID, response.Shipping{}
	}
	err, res := as.shippingRepository.GetShipping(id)
	if err != nil {
		return err, response.Shipping{}
	}
	return nil, res
}

func (as *ShippingService) UpdateShipping(id string, data request.Shipping) (error, response.Shipping) {
	if id == "" {
		return errors.ERR_GET_SHIPPING_BAD_REQUEST_ID, response.Shipping{}
	}
	err, res := as.shippingRepository.UpdateShipping(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.Shipping{}
	}
	return nil, res
}

func (as *ShippingService) DeleteShipping(id string) (error, response.Shipping) {
	if id == "" {
		return errors.ERR_GET_SHIPPING_BAD_REQUEST_ID, response.Shipping{}
	}
	err, res := as.shippingRepository.DeleteShipping(id)

	if err != nil {
		return errors.ERR_DELETE_SHIPPING, response.Shipping{}
	}

	return nil, res
}
