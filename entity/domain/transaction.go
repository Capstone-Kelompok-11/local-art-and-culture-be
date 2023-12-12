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

func ConvertModelTransactionsToResponse(transactions []models.Transaction) []response.Transaction {
	var responseTransactions []response.Transaction

	for _, modelTransaction := range transactions {
		responseTransaction := ConvertFromModelToTransactionRes(modelTransaction)
		responseTransactions = append(responseTransactions, *responseTransaction)
	}

	return responseTransactions
}
