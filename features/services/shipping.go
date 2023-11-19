package services

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"

	// "lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"
	// "lokasani/helpers/middleware"
)

type IShippingService interface {
	CreateShippingMethod(data *request.Shipping) (response.Shipping, error)
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

func (sms *ShippingService) GetAllShipping() (error, []response.Shipping) {
	err, res := sms.shippingRepository.GetAllShipping()
	if err != nil {
		return err, nil
	}
	return nil, res
}

func (sms *ShippingService) GetShipping(id string) (error, response.Shipping) {
	if id == "" {
		return errors.ERR_GET_SHIPPING_BAD_REQUEST_ID, response.Shipping{}
	}
	err, res := sms.shippingRepository.GetShipping(id)
	if err != nil {
		return err, response.Shipping{}
	}
	return nil, res
}

func (sms *ShippingService) UpdateShipping(id string, data request.Shipping) (error, response.Shipping) {
	if id == "" {
		return errors.ERR_GET_SHIPPING_BAD_REQUEST_ID, response.Shipping{}
	}
	err, res := sms.shippingRepository.UpdateShipping(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.Shipping{}
	}
	return nil, res
}

func (sms *ShippingService) DeleteShipping(id string) (error, response.Shipping) {
	if id == "" {
		return errors.ERR_GET_SHIPPING_BAD_REQUEST_ID, response.Shipping{}
	}
	err, res := sms.shippingRepository.DeleteShipping(id)
	fmt.Println(res)
	if err != nil {
		return errors.ERR_DELETE_SHIPPING, response.Shipping{}
	}

	return nil, res
}
