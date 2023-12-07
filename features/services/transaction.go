package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
	"lokasani/helpers/midtrans"
)

type ITransactionService interface {
	CreateTransaction(data *request.Transaction) (response.Transaction, error)
	GetAllTransaction() ([]response.Transaction, error)
	GetTransaction(id string) (response.Transaction, error)
	UpdateTransaction(id string, data request.Transaction) (response.Transaction, error)
	DeleteTransaction(id string) (response.Transaction, error)
}

type TransactionService struct {
	transactionRepository       repositories.ITransactionRepository
	transactionDetailRepository repositories.ITransactionDetailRepository
}

func NewTransactionService(repo repositories.ITransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: repo}
}

func (rs *TransactionService) CreateTransaction(data *request.Transaction) (response.Transaction, error) {
	if data.TransactionDate.IsZero() {
		return response.Transaction{}, errors.ERR_EVENT_DATE_IS_EMPTY
	}

	res, err := rs.transactionRepository.CreateTransaction(data)
	if err != nil {
		return response.Transaction{}, errors.ERR_CREATE_TRANSACTION_DATABASE
	}
	// fmt.Println(data.TransactionDetail)

	// var transactionDetailRes []response.TransactionDetail
	// for i := range data.TransactionDetail {
	// 	data.TransactionDetail[i].TransactionId = res.Id
	// 	result, err := rs.transactionDetailRepository.CreateTransactionDetail(&data.TransactionDetail[i])
	// 	if err != nil {
	// 		return res, errors.ERR_CREATE_TRANSACTION_DETAIL
	// 	}
	// 	transactionDetailRes = append(transactionDetailRes, result)
	// }
	// res.TransactionDetail = transactionDetailRes
	res.Total = 10000
	midtrans.Payment(&res)

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
	var transactionDetailRes []response.TransactionDetail
	for i := range data.TransactionDetail {
		data.TransactionDetail[i].TransactionId = res.Id
		err, result := rs.transactionDetailRepository.UpdateTransactionDetail(string(rune(data.TransactionDetail[i].Id)), data.TransactionDetail[i])
		if err != nil {
			return res, errors.ERR_CREATE_TRANSACTION_DETAIL
		}
		transactionDetailRes = append(transactionDetailRes, result)
	}
	res.TransactionDetail = transactionDetailRes
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
