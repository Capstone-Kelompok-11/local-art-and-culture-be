package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromTransactionReqToModel(data request.Transaction) *models.Transaction {
	return &models.Transaction{
		Model: gorm.Model{
			ID: data.Id,
		},
		TransactionDate:   data.TransactionDate,
		UserId:            data.UserId,
		PaymentMethodId:   data.PaymentMethodId,
		ShippingMethodId:  data.ShippingMethodId,
		Status:            data.Status,
		TransactionNumber: data.TransactionNumber,
		User:              *ConvertFromUserReqToModel(data.User),
		Payment:           *ConvertFromPaymentReqToModel(data.Payment),
		Shipping:          *ConvertFromShippingReqToModel(data.Shipping),
		TransactionDetail: *ConvertFromTransactionDetailReqToModels(data.TransactionDetail),
	}
}

func ConvertFromModelToTransactionRes(data models.Transaction) *response.Transaction {
	return &response.Transaction{
		Id:                data.ID,
		TransactionDate:   data.TransactionDate,
		UserId:            data.UserId,
		PaymentMethodId:   data.PaymentMethodId,
		ShippingMethodId:  data.ShippingMethodId,
		Status:            data.Status,
		TransactionNumber: data.TransactionNumber,
		User:              *ConvertFromModelToUserRes(data.User),
		Payment:           *ConvertFromModelToPaymentRes(data.Payment),
		Shipping:          *ConvertFromModelToShippingRes(data.Shipping),
		TransactionDetail: *ConvertFromModelsToTransactionDetailRes(data.TransactionDetail),
	}
}

func ConvertModelTransactionsToResponse(transactions []models.Transaction) []*response.Transaction {
	list := []*response.Transaction{}

	for _, v := range transactions {
		res:= ConvertFromModelToTransactionRes(v)
		list = append(list, res)
	}
	return list
}

func ConvertFromModelToTransactionReport(data models.TransactionReport) *response.TransactionReport {
	return &response.TransactionReport{
		Id: 			 data.Id,
		TransactionDate: data.TransactionDate,
		Status: 		 data.Status,
		Qty: 			 data.Qty,
		Price: 			 data.Price,
		Nominal: 		 data.Nominal,
	}
}