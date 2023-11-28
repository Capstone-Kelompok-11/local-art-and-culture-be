package repositories

import (
	"fmt"
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ITransactionDetailRepository interface {
	CreateTransactionDetail(data *request.TransactionDetail) (response.TransactionDetail, error)
	GetAllTransactionDetail() (error, []response.TransactionDetail)
	GetTransactionDetail(id string) (error, response.TransactionDetail)
	UpdateTransactionDetail(id string, input request.TransactionDetail) (error, response.TransactionDetail)
	DeleteTransactionDetail(id string) (error, response.TransactionDetail)
}

type transactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) *transactionDetailRepository {
	return &transactionDetailRepository{db}
}

func (ar *transactionDetailRepository) CreateTransactionDetail(data *request.TransactionDetail) (response.TransactionDetail, error) {
	dataTransactionDetail := domain.ConvertFromTransactionDetailReqToModel(*data)

	err := ar.db.Create(&dataTransactionDetail).Error
	if err != nil {
		return response.TransactionDetail{}, err
	}
	fmt.Println(data)
	return *domain.ConvertFromModelToTransactionDetailRes(*dataTransactionDetail), nil
}

func (ar *transactionDetailRepository) GetAllTransactionDetail() (error, []response.TransactionDetail) {
	var allTransactionDetail []models.TransactionDetail
	var resAllTransactionDetail []response.TransactionDetail
	err := ar.db.Find(&allTransactionDetail).Error
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}

	for i := 0; i < len(allTransactionDetail); i++ {
		transactionDetailVm := domain.ConvertFromModelToTransactionDetailRes(allTransactionDetail[i])
		resAllTransactionDetail = append(resAllTransactionDetail, *transactionDetailVm)
	}
	return nil, resAllTransactionDetail
}

func (ar *transactionDetailRepository) GetTransactionDetail(id string) (error, response.TransactionDetail) {
	var transactionDetailData models.TransactionDetail
	err := ar.db.First(&transactionDetailData, "id = ?", id).Error

	if err != nil {
		return err, response.TransactionDetail{}
	}

	return nil, *domain.ConvertFromModelToTransactionDetailRes(transactionDetailData)
}

func (ar *transactionDetailRepository) UpdateTransactionDetail(id string, input request.TransactionDetail) (error, response.TransactionDetail) {
	transactionDetailData := models.TransactionDetail{}
	err := ar.db.First(&transactionDetailData, "id = ?", id).Error

	if input.Qty != 0 {
		transactionDetailData.Qty = input.Qty
	}
	if input.ProductId != nil {
		transactionDetailData.ProductId = input.ProductId
	}
	if input.TicketId != nil {
		transactionDetailData.TicketId = input.TicketId
	}
	if input.Fullname != nil {
		transactionDetailData.Fullname = input.Fullname
	}
	if input.Contact != nil {
		transactionDetailData.Contact = input.Contact
	}
	if input.Email != nil {
		transactionDetailData.Email = input.Email
	}

	if err != nil {
		return err, response.TransactionDetail{}
	}

	if err = ar.db.Save(&transactionDetailData).Error; err != nil {
		return err, response.TransactionDetail{}
	}
	return nil, *domain.ConvertFromModelToTransactionDetailRes(transactionDetailData)
}

func (ar *transactionDetailRepository) DeleteTransactionDetail(id string) (error, response.TransactionDetail) {
	transactionDetailData := models.TransactionDetail{}
	res := response.TransactionDetail{}
	find := ar.db.First(&transactionDetailData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToTransactionDetailRes(transactionDetailData)
	}
	err := ar.db.Delete(&transactionDetailData, "id = ?", id).Error
	if err != nil {
		return err, response.TransactionDetail{}
	}
	return nil, res
}
