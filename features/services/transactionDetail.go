package services

// import (
// 	"lokasani/entity/request"
// 	"lokasani/entity/response"
// 	"lokasani/features/repositories"
// 	"lokasani/helpers/errors"
// )

// type ITransactionDetailService interface {
// 	CreateTransactionDetail(data *request.TransactionDetail) (response.TransactionDetail, error)
// 	GetAllTransactionDetail(nameFilter string) ([]response.TransactionDetail, error)
// 	GetTransactionDetail(id string) (response.TransactionDetail, error)
// 	UpdateTransactionDetail(id string, input request.TransactionDetail) (response.TransactionDetail, error)
// 	DeleteTransactionDetail(id string) (response.TransactionDetail, error)
// }

// type TransactionDetailService struct {
// 	transactionDetailRepository repositories.ITransactionDetailRepository
// }

// func NewTransactionDetailService(repo repositories.ITransactionDetailRepository) *TransactionDetailService {
// 	return &TransactionDetailService{transactionDetailRepository: repo}
// }

// func (ti *TransactionDetailService) CreateTransactionDetail(data *request.TransactionDetail) (response.TransactionDetail, error) {
// 	if data.Type == "" {
// 		return response.TransactionDetail{}, errors.ERR_TYPE_IS_EMPTY
// 	}
// 	if data.Price == 0 {
// 		return response.TransactionDetail{}, errors.ERR_PRICE_IS_EMPTY
// 	}
// 	if data.StartTime.IsZero() || data.EndTime.IsZero() {
// 		return response.TransactionDetail{}, errors.ERR_EVENT_DATE_IS_EMPTY
// 	}
// 	if data.Qty == 0 {
// 		return response.TransactionDetail{}, errors.ERR_QTY_IS_EMPTY
// 	}
// 	if data.Name == "" {
// 		return response.TransactionDetail{}, errors.ERR_NAME_IS_EMPTY
// 	}
// 	if data.Description == "" {
// 		return response.TransactionDetail{}, errors.ERR_DESCRIPTION_IS_EMPTY
// 	}

// 	res, err := ti.transactionDetailRepository.CreateTransactionDetail(data)
// 	if err != nil {
// 		return response.TransactionDetail{}, errors.ERR_CREATE_TransactionDetail_DATABASE
// 	}
// 	return res, nil
// }

// func (ti *TransactionDetailService) GetAllTransactionDetail() ([]response.TransactionDetail, error) {
// 	res, err := ti.transactionDetailRepository.GetAllTransactionDetail()
// 	if err != nil {
// 		return nil, errors.ERR_GET_DATA
// 	}
// 	return res, nil
// }

// func (ti *TransactionDetailService) GetTransactionDetail(id string) (response.TransactionDetail, error) {
// 	if id == "" {
// 		return response.TransactionDetail{}, errors.ERR_GET_TransactionDetail_BAD_REQUEST_ID
// 	}
// 	err,res := ti.transactionDetailRepository.GetTransactionDetail(id)
// 	if err != nil {
// 		return response.TransactionDetail{}, err
// 	}
// 	return res, nil
// }

// func (ti *TransactionDetailService) UpdateTransactionDetail(id string, input request.TransactionDetail) (response.TransactionDetail, error) {
// 	if id == "" {
// 		return response.TransactionDetail{}, errors.ERR_GET_TransactionDetail_BAD_REQUEST_ID
// 	}
// 	res, err := ti.transactionDetailRepository.UpdateTransactionDetail(id, input)
// 	if err != nil {
// 		return response.TransactionDetail{}, err
// 	}
// 	return res, nil
// }

// func (ti *TransactionDetailService) DeleteTransactionDetail(id string) (response.TransactionDetail, error) {
// 	if id == "" {
// 		return response.TransactionDetail{}, errors.ERR_GET_TransactionDetail_BAD_REQUEST_ID
// 	}
// 	res, err := ti.transactionDetailRepository.DeleteTransactionDetail(id)
// 	if err != nil {
// 		return response.TransactionDetail{}, err
// 	}
// 	return res, nil
// }
