package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ITransactionService interface {
	CreateTransaction(data *request.Transaction) (response.Transaction, error)
	GetAllTransaction() ([]response.Transaction, error)
	GetTransaction(id string) (response.Transaction, error)
	UpdateTransaction(id string, data request.Transaction) (response.Transaction, error)
	DeleteTransaction(id string) (response.Transaction, error)
}

type TransactionService struct {
	transactionRepository repositories.ITransactionRepository
}

func NewTransactionService(repo repositories.ITransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: repo}
}

func (rs *TransactionService) CreateTransaction(data *request.Transaction) (response.Transaction, error) {
	if data.Status == "" {
		return response.Transaction{}, errors.ERR_STATUS_IS_EMPTY
	}
	if data.TransactionDate.IsZero() {
		return response.Transaction{}, errors.ERR_EVENT_DATE_IS_EMPTY
	}

	res, err := rs.transactionRepository.CreateTransaction(data)
	if err != nil {
		return response.Transaction{}, errors.ERR_CREATE_TRANSACTION_DATABASE
	}

	return res, nil
}

func (rs *TransactionService) GetAllTransaction() ([]response.Transaction, error) {
	err, res := rs.transactionRepository.GetAllTransaction()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (rs *TransactionService) GetTransaction(id string) (response.Transaction, error) {
	if id == "" {
		return response.Transaction{}, errors.ERR_GET_TRANSACTION_BAD_REQUEST_ID
	}
	err, res := rs.transactionRepository.GetTransaction(id)
	if err != nil {
		return response.Transaction{}, errors.ERR_GET_DATA
	}
	return res, nil
}

func (rs *TransactionService) UpdateTransaction(id string, data request.Transaction) (response.Transaction, error) {
	if id == "" {
		return response.Transaction{}, errors.ERR_GET_TRANSACTION_BAD_REQUEST_ID
	}
	err, res := rs.transactionRepository.UpdateTransaction(id, data)
	if err != nil {
		return response.Transaction{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (rs *TransactionService) DeleteTransaction(id string) (response.Transaction, error) {
	if id == "" {
		return response.Transaction{}, errors.ERR_GET_TRANSACTION_BAD_REQUEST_ID
	}
	err, res := rs.transactionRepository.DeleteTransaction(id)

	if err != nil {
		return response.Transaction{}, errors.ERR_DELETE_TRANSACTION
	}

	return res, nil
}
