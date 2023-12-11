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
	GetHistoryTransaction(userID uint, page, pageSize int) ([]*response.Transaction, int, error)
	GetReportTransaction(creatorId uint, role string) ([]response.Transaction, error)
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

func (ar *transactionRepository) GetHistoryTransaction(userID uint, page, pageSize int) ([]*response.Transaction, int, error) {
	var transactions []models.Transaction

	query := ar.db.Where("user_id = ?", userID)

	var count int64
	if err := query.Model(&models.Transaction{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	query = query.Limit(pageSize).Offset(offset).
		Preload("TransactionDetail").
		Preload("TransactionDetail.Product").
		Preload("TransactionDetail.Product.Creator").
		Preload("TransactionDetail.Product.Category").
		Preload("TransactionDetail.Ticket").
		Preload("TransactionDetail.Ticket.Event").
		Preload("User").
		Preload("User.Role").
		Preload("Shipping").
		Preload("Payment")

	if err := query.Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	resAllHistory := domain.ConvertModelTransactionsToResponse(transactions)

	return resAllHistory, int(count), nil
}

func (ar *transactionRepository) GetReportTransaction(creatorId uint, role string) ([]response.Transaction, error) {
	var allTransaction []models.Transaction
	var resAllTransaction []response.Transaction

	query := ar.db.Preload("TransactionDetail")
	if role == consts.ProductCreator {
		fmt.Println("s"+role)
		query.Preload("TransactionDetail.Product", "creator_id = ? ", creatorId).Preload("TransactionDetail.Product.Creator").
		Preload("TransactionDetail.Product.Category")
	} else if role == consts.EventCreator{
		fmt.Println("s"+role)
		query.Preload("TransactionDetail.Ticket.Event", "creator_id = ? ", creatorId).Preload("TransactionDetail.Ticket.Event")
	}

	err := query.Find(&allTransaction).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allTransaction); i++ {
		transactionVm := domain.ConvertFromModelToTransactionRes(allTransaction[i])
		resAllTransaction = append(resAllTransaction, *transactionVm)
	}
	return resAllTransaction, nil
}