package repositories

import (
	"fmt"
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"time"

	consts "lokasani/helpers/const"
	"lokasani/helpers/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(data *request.Transaction) (response.Transaction, error)
	GetAllTransaction() (error, []response.Transaction)
	GetTransaction(id string) (error, response.Transaction)
	UpdateTransaction(id string, input request.Transaction) (error, response.Transaction)
	DeleteTransaction(id string) (error, response.Transaction)
	GetTransactionReport(userID uint) ([]*response.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (ar *transactionRepository) CreateTransaction(data *request.Transaction) (response.Transaction, error) {
	newUUID := uuid.New()
	uuidString := newUUID.String()
	dataTransaction := domain.ConvertFromTransactionReqToModel(*data)
	dataTransaction.TransactionDate = time.Now()
	dataTransaction.Status = consts.OrderStatusUnpaid
	dataTransaction.TransactionNumber = "LKSNI/" + uuidString

	if dataTransaction.PaymentMethodId == 0 {
		dataTransaction.PaymentMethodId = *getPaymentMethod(ar.db, data.PaymentMethodId, "payment")
	}
	// if dataTransaction.ShippingMethodId == nil {
	// 	dataTransaction.ShippingMethodId = getPaymentMethod(ar.db, *data.ShippingMethodId, "shipping")
	// }
	err := ar.db.Create(&dataTransaction).Error
	if err != nil {
		return response.Transaction{}, err
	}
	return *domain.ConvertFromModelToTransactionRes(*dataTransaction), nil
}

func getPaymentMethod(db *gorm.DB, id uint, from string) *uint {
	if from == "payment" {
		var paymentMethod models.Payment
		if id != 0 {
			err := db.First(&paymentMethod, "id = ?", id).Error
			if err != nil {
				err = db.First(&paymentMethod).Error
				return &paymentMethod.ID
			}
			return &paymentMethod.ID
		} else {
			db.First(&paymentMethod)
			return &paymentMethod.ID
		}
	} else {
		var shippingMethod models.Shipping
		if id != 0 {
			err := db.First(&shippingMethod, "id = ?", id).Error
			if err != nil {
				err = db.First(&shippingMethod).Error
				return &shippingMethod.ID
			}
			return &shippingMethod.ID
		} else {
			db.First(&shippingMethod)
			return &shippingMethod.ID
		}
	}
}

func (ar *transactionRepository) GetAllTransaction() (error, []response.Transaction) {
	var allTransaction []models.Transaction
	var resAllTransaction []response.Transaction
	err := ar.db.Preload("TransactionDetail").Find(&allTransaction).Error
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}

	for i := 0; i < len(allTransaction); i++ {
		transactionVm := domain.ConvertFromModelToTransactionRes(allTransaction[i])
		resAllTransaction = append(resAllTransaction, *transactionVm)
	}
	return nil, resAllTransaction
}

func (ar *transactionRepository) GetTransaction(id string) (error, response.Transaction) {
	var transactionData models.Transaction
	err := ar.db.First(&transactionData, "id = ?", id).Error

	if err != nil {
		return err, response.Transaction{}
	}

	return nil, *domain.ConvertFromModelToTransactionRes(transactionData)
}

func (ar *transactionRepository) UpdateTransaction(id string, input request.Transaction) (error, response.Transaction) {
	transactionData := models.Transaction{}
	fmt.Println("repo transaksi" + input.TransactionNumber)
	err := ar.db.First(&transactionData, func(db *gorm.DB) {
		if input.TransactionNumber != "" {
			fmt.Println(2)
			db.Where("transaction_number = ?", input.TransactionNumber)
		} else {
			fmt.Println(2)
			db.Where("id = ? ", id)
		}
	}).Error

	if err != nil {
		return err, response.Transaction{}
	}

	if input.Status != "" {
		transactionData.Status = input.Status
	}
	if input.PaymentMethodId != 0 {
		transactionData.PaymentMethodId = input.PaymentMethodId
	}
	if input.ShippingMethodId != nil {
		transactionData.ShippingMethodId = input.ShippingMethodId
	}
	if !input.TransactionDate.IsZero() {
		transactionData.TransactionDate = input.TransactionDate
	}

	if err = ar.db.Save(&transactionData).Error; err != nil {
		return err, response.Transaction{}
	}
	return nil, *domain.ConvertFromModelToTransactionRes(transactionData)
}

func (ar *transactionRepository) DeleteTransaction(id string) (error, response.Transaction) {
	transactionData := models.Transaction{}
	res := response.Transaction{}
	find := ar.db.First(&transactionData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToTransactionRes(transactionData)
	}
	err := ar.db.Delete(&transactionData, "id = ?", id).Error
	if err != nil {
		return err, response.Transaction{}
	}
	return nil, res
}

func (ar *transactionRepository) GetTransactionReport(userID uint) ([]*response.Transaction, error) {
	var transactions []models.Transaction
	err := ar.db.
		Preload("TransactionDetail").Preload("TransactionDetail.Product").Preload("TransactionDetail.Product.Creator").
		Preload("TransactionDetail.Product.Category").Preload("TransactionDetail.Ticket").Preload("TransactionDetail.Ticket.Event").
		Preload("User").Preload("User.Role").Preload("Shipping").Preload("Payment").Find(&transactions, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	resAllHistory := domain.ConvertModelTransactionsToResponse(transactions)

	return resAllHistory, nil
}