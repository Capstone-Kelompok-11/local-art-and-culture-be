package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ITransactionDetailService interface {
	CreateTransactionDetail(data *request.TransactionDetail) (response.TransactionDetail, error)
	GetAllTransactionDetail() ([]response.TransactionDetail, error)
	GetTransactionDetail(id string) (response.TransactionDetail, error)
	UpdateTransactionDetail(id string, input request.TransactionDetail) (response.TransactionDetail, error)
	DeleteTransactionDetail(id string) (response.TransactionDetail, error)
}

type TransactionDetailService struct {
	transactionDetailRepository repositories.ITransactionDetailRepository
}

func NewTransactionDetailService(repo repositories.ITransactionDetailRepository) *TransactionDetailService {
	return &TransactionDetailService{transactionDetailRepository: repo}
}

func (tds *TransactionDetailService) CreateTransactionDetail(data *request.TransactionDetail) (response.TransactionDetail, error) {

	res, err := tds.transactionDetailRepository.CreateTransactionDetail(data)
	if err != nil {
		return response.TransactionDetail{}, errors.ERR_CREATE_TRANSACTION_DETAIL
	}
	return res, nil
}

func (tds *TransactionDetailService) GetAllTransactionDetail() ([]response.TransactionDetail, error) {
	err, res := tds.transactionDetailRepository.GetAllTransactionDetail()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (tds *TransactionDetailService) GetTransactionDetail(id string) (response.TransactionDetail, error) {
	if id == "" {
		return response.TransactionDetail{}, errors.ERR_GET_TRANSACTION_DETAIL_BAD_REQUEST_ID
	}
	err, res := tds.transactionDetailRepository.GetTransactionDetail(id)
	if err != nil {
		return response.TransactionDetail{}, err
	}
	return res, nil
}

func (tds *TransactionDetailService) UpdateTransactionDetail(id string, input request.TransactionDetail) (response.TransactionDetail, error) {
	if id == "" {
		return response.TransactionDetail{}, errors.ERR_GET_TRANSACTION_DETAIL_BAD_REQUEST_ID
	}
	err, res := tds.transactionDetailRepository.UpdateTransactionDetail(id, input)
	if err != nil {
		return response.TransactionDetail{}, err
	}
	return res, nil
}

func (tds *TransactionDetailService) DeleteTransactionDetail(id string) (response.TransactionDetail, error) {
	if id == "" {
		return response.TransactionDetail{}, errors.ERR_GET_TRANSACTION_DETAIL_BAD_REQUEST_ID
	}
	err, res := tds.transactionDetailRepository.DeleteTransactionDetail(id)
	if err != nil {
		return response.TransactionDetail{}, err
	}
	return res, nil
}
