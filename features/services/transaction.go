package services

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
	"lokasani/helpers/midtrans"
	"math"
)

type ITransactionService interface {
	CreateTransaction(data *request.Transaction) (response.Transaction, error)
	ConfirmPayment(id, fraudStatus string) (response.Transaction, error)
	GetAllTransaction() ([]response.Transaction, error)
	GetTransaction(id string) (response.Transaction, error)
	UpdateTransaction(id string, data request.Transaction) (response.Transaction, error)
	DeleteTransaction(id string) (response.Transaction, error)
	GetHistoryTransaction(userID uint, page, pageSize int) ([]*response.Transaction, int, error)
	GetReportTransaction(creatorId uint, role string) ([]response.TransactionReport, error)
	CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	GetNextPage(currentPage, allPages int) int
	GetPrevPage(currentPage int) int
}

type TransactionService struct {
	transactionRepository       repositories.ITransactionRepository
	transactionDetailRepository repositories.ITransactionDetailRepository
	midtransService             midtrans.IMidtransService
}

func NewTransactionService(repo repositories.ITransactionRepository, TDRepo repositories.ITransactionDetailRepository, midtransService midtrans.IMidtransService) *TransactionService {
	return &TransactionService{
		transactionRepository:       repo,
		transactionDetailRepository: TDRepo,
		midtransService:             midtransService,
	}
}

func (rs *TransactionService) CreateTransaction(data *request.Transaction) (response.Transaction, error) {
	if data.TransactionDetail == nil {
		return response.Transaction{}, errors.ERR_TRANSACTION_DETAIL_EMPTY
	}
	res, err := rs.transactionRepository.CreateTransaction(data)
	if err != nil {
		return response.Transaction{}, errors.ERR_CREATE_TRANSACTION_DATABASE
	}

	var transactionDetailRes []response.TransactionDetail
	for i := range data.TransactionDetail {
		data.TransactionDetail[i].TransactionId = res.Id
		result, err := rs.transactionDetailRepository.CreateTransactionDetail(&data.TransactionDetail[i])
		if err != nil {
			return res, errors.ERR_CREATE_TRANSACTION_DETAIL
		}
		if data.TransactionDetail[i].ProductId != nil {
			result.Subtotal = int64(data.TransactionDetail[i].Qty) * int64(result.Product.Price)
		} else {
			result.Subtotal = int64(data.TransactionDetail[i].Qty) * int64(data.TransactionDetail[i].Ticket.Price)
		}

		transactionDetailRes = append(transactionDetailRes, result)
	}
	res.TransactionDetail = transactionDetailRes
	fmt.Println(transactionDetailRes)
	res.Total = sumTotal(res.TransactionDetail)
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

func (rs *TransactionService) ConfirmPayment(id, fraudStatus string) (response.Transaction, error) {
	if id == "" {
		return response.Transaction{}, errors.ERR_GET_TRANSACTION_BAD_REQUEST_ID
	}
	fmt.Println(id)
	err, res := rs.midtransService.Verification(id, fraudStatus)

	if err != nil {
		return response.Transaction{}, err
	}

	return res, nil
}

func (rs *TransactionService) GetHistoryTransaction(userID uint, page, pageSize int) ([]*response.Transaction, int, error) {
	if userID == 0 {
		return []*response.Transaction{}, 0, errors.ERR_GET_TRANSACTION_BAD_REQUEST_ID
	}

	res, transactions, err := rs.transactionRepository.GetHistoryTransaction(userID, page, pageSize)
	if err != nil {
		return []*response.Transaction{}, 0, errors.ERR_GET_DATA
	}

	return res, transactions, nil
}

func (rs *TransactionService) GetReportTransaction(creatorId uint, role string) ([]response.TransactionReport, error) {
	res, err := rs.transactionRepository.GetReportTransaction(creatorId, role)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (rs *TransactionService) CalculatePaginationValues(page, pageSize, allItmes int) (int, int) {
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

func (rs *TransactionService) GetNextPage(currentPage, allPages int) int {
	if currentPage < allPages {
		return currentPage + 1
	}
	return allPages
}

func (rs *TransactionService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}

func sumTotal(input []response.TransactionDetail) int64 {
	var total int64
	for i := range input {
		total += input[i].Subtotal
	}
	return total
}
