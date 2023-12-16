package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IShippingRepository interface {
	CreateShippingMethod(data *request.Shipping) (response.Shipping, error)
	GetAllShipping() ([]response.Shipping, error)
	GetShipping(id string) (response.Shipping, error)
	UpdateShipping(id string, input request.Shipping) (response.Shipping, error)
	DeleteShipping(id string) (response.Shipping, error)
}

type shippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) *shippingRepository {
	return &shippingRepository{db}
}

func (ar *shippingRepository) CreateShippingMethod(data *request.Shipping) (response.Shipping, error) {
	dataShipping := domain.ConvertFromShippingReqToModel(*data)
	err := ar.db.Create(&dataShipping).Error
	if err != nil {
		return response.Shipping{}, err
	}
	return *domain.ConvertFromModelToShippingRes(*dataShipping), nil
}

func (ar *shippingRepository) GetAllShipping() ([]response.Shipping, error) {
	var allShipping []models.Shipping
	var resAllShipping []response.Shipping
	err := ar.db.Find(&allShipping).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allShipping); i++ {
		shippingVm := domain.ConvertFromModelToShippingRes(allShipping[i])
		resAllShipping = append(resAllShipping, *shippingVm)
	}
	return resAllShipping, nil
}

func (ar *shippingRepository) GetShipping(id string) (response.Shipping, error) {
	var shippingData models.Shipping
	err := ar.db.First(&shippingData, "id = ?", id).Error

	if err != nil {
		return response.Shipping{}, err
	}

	return *domain.ConvertFromModelToShippingRes(shippingData), nil
}

func (ar *shippingRepository) UpdateShipping(id string, input request.Shipping) (response.Shipping, error) {
	shippingData := models.Shipping{}
	err := ar.db.First(&shippingData, "id = ?", id).Error

	if err != nil {
		return response.Shipping{}, err
	}

	if input.Name != "" {
		shippingData.Name = input.Name
	}

	if err = ar.db.Save(&shippingData).Error; err != nil {
		return response.Shipping{}, err
	}
	return *domain.ConvertFromModelToShippingRes(shippingData), nil
}

func (ar *shippingRepository) DeleteShipping(id string) (response.Shipping, error) {
	shippingData := models.Shipping{}
	res := response.Shipping{}
	find := ar.db.First(&shippingData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToShippingRes(shippingData)
	}
	err := ar.db.Delete(&shippingData, "id = ?", id).Error
	if err != nil {
		return response.Shipping{}, err
	}
	return res, nil
}