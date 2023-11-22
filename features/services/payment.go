package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IPaymentService interface {
	CreatePayment(data *request.Payment) (response.Payment, error)
	GetAllPayment() ([]response.Payment, error)
	GetPayment(id string) (response.Payment, error)
	UpdatePayment(id string, input request.Payment) (response.Payment, error)
	DeletePayment(id string) (response.Payment, error)
}

type PaymentService struct {
	paymentRepository repositories.IPaymentRepository
}

func NewPaymentService(repo repositories.IPaymentRepository) *PaymentService {
	return &PaymentService{paymentRepository: repo}
}

func (pa *PaymentService) CreatePayment(data *request.Payment) (response.Payment, error) {
	if data.Name == "" {
		return response.Payment{}, errors.ERR_NAME_IS_EMPTY
	}

	res, err := pa.paymentRepository.CreatePayment(data)
	if err != nil {
		return response.Payment{}, errors.ERR_CREATE_PAYMENT
	}
	return res, nil
}

func (pa *PaymentService) GetAllPayment() ([]response.Payment, error) {
	err, res := pa.paymentRepository.GetAllPayment()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return nil, res
}

func (pa *PaymentService) GetPayment(id string) (response.Payment, error) {
	if id == "" {
		return response.Payment{}, errors.ERR_GET_PAYMENT_BAD_REQUEST_ID
	}

	res, err := pa.paymentRepository.GetPayment(id)
	if err != nil {
		return response.Payment{}, errors.ERR_GET_DATA
	}
	return res, nil
}

func (pa *PaymentService) UpdatePayment(id string, input request.Payment) (response.Payment, error) {
	if id == "" {
		return response.Payment{}, errors.ERR_GET_PAYMENT_BAD_REQUEST_ID
	}
	res, err := pa.paymentRepository.UpdatePayment(id, input)
	if err != nil {
		return response.Payment{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (pa *PaymentService) DeletePayment(id string) (response.Payment, error) {
	if id == "" {
		return response.Payment{}, errors.ERR_GET_PAYMENT_BAD_REQUEST_ID
	}

	res, err := pa.paymentRepository.DeletePayment(id)
	if err != nil {
		return response.Payment{}, errors.ERR_DELETE
	}
	return res, nil
}