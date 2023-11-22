package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IPaymentRepository interface {
	CreatePayment(data *request.Payment) (response.Payment, error) 
	GetAllPayment() ([]response.Payment, error)
	GetPayment(id string) (response.Payment, error)
	UpdatePayment(id string, input request.Payment) (response.Payment, error)
	DeletePayment(id string) (response.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (pa *paymentRepository) CreatePayment(data *request.Payment) (response.Payment, error) {
	dataPayment := domain.ConvertFromPaymentReqToModel(*data)
	err := pa.db.Create(&dataPayment).Error
	if err != nil {
		return response.Payment{}, err
	}
	return *domain.ConvertFromModelToPaymentRes(*dataPayment), nil
}

func (pa *paymentRepository) GetAllPayment() ([]response.Payment, error) {
	var allPayment []models.Payment
	var resAllPayment []response.Payment
	err := pa.db.Find(&allPayment).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allPayment); i++ {
		paymentVm := domain.ConvertFromModelToPaymentRes(allPayment[i])
		resAllPayment = append(resAllPayment, *paymentVm)
	}
	return resAllPayment, nil
}

func (pa *paymentRepository) GetPayment(id string) (response.Payment, error) {
	var paymentData models.Payment
	err := pa.db.First(&paymentData, "id = ?", id).Error
	if err != nil {
		return response.Payment{}, err
	}
	return *domain.ConvertFromModelToPaymentRes(paymentData), nil
}

func (pa *paymentRepository) UpdatePayment(id string, input request.Payment) (response.Payment, error) {
	paymentData := models.Payment{}
	err := pa.db.First(&paymentData, "id = ?", id).Error
	if err != nil {
		return response.Payment{}, errors.ERR_GET_PAYMENT_BAD_REQUEST_ID
	}

	if input.Name != "" {
		paymentData.Name = input.Name
	}

	if err = pa.db.Save(&paymentData).Error; err != nil{
		return response.Payment{}, errors.ERR_UPDATE_DATA
	}
	return *domain.ConvertFromModelToPaymentRes(paymentData), nil
}

func (pa *paymentRepository) DeletePayment(id string) (response.Payment, error) {
	paymentData := models.Payment{}
	res := response.Payment{}
	find := pa.db.First(&paymentData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToPaymentRes(paymentData)
	}
	err := pa.db.Delete(&paymentData, "id = ?", id).Error
	if err != nil {
		return response.Payment{}, err
	}
	return res, nil
}