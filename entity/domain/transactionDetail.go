package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromTransactionDetailReqToModels(data []request.TransactionDetail) *[]models.TransactionDetail {
	var result []models.TransactionDetail
	var temp models.TransactionDetail
	for i := range data {
		temp.ID = data[i].Id
		temp.TransactionId = data[i].TransactionId
		temp.ProductId = data[i].ProductId
		temp.TicketId = data[i].TicketId
		temp.Qty = data[i].Qty
		temp.Fullname = data[i].Fullname
		temp.Contact = data[i].Contact
		result = append(result, temp)
	}
	return &result
}

func ConvertFromTransactionDetailReqToModel(data request.TransactionDetail) *models.TransactionDetail {
	return &models.TransactionDetail{
		Model: gorm.Model{
			ID: data.Id,
		},
		TransactionId: data.TransactionId,
		ProductId:     data.ProductId,
		TicketId:      data.TicketId,
		Qty:           data.Qty,
		Fullname:      data.Fullname,
		Contact:       data.Contact,
	}
}

func ConvertFromModelsToTransactionDetailRes(data []models.TransactionDetail) *[]response.TransactionDetail {
	var result []response.TransactionDetail
	var temp response.TransactionDetail
	for i := range data {
		temp.Id = data[i].ID
		temp.TransactionId = data[i].TransactionId
		temp.ProductId = data[i].ProductId
		temp.TicketId = data[i].TicketId
		temp.Qty = data[i].Qty
		temp.Fullname = data[i].Fullname
		temp.Contact = data[i].Contact
		result = append(result, temp)
	}
	return &result
}

func ConvertFromModelToTransactionDetailRes(data models.TransactionDetail) *response.TransactionDetail {
	return &response.TransactionDetail{
		Id:            data.ID,
		TransactionId: data.TransactionId,
		ProductId:     data.ProductId,
		TicketId:      data.TicketId,
		Qty:           data.Qty,
		Fullname:      data.Fullname,
		Contact:       data.Contact,
	}
}
