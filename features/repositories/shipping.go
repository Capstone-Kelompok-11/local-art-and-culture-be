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

type IShippingRepository interface {
	CreateShippingMethod(data *request.Shipping) (response.Shipping, error)
	GetAllShipping() (error, []response.Shipping)
	GetShipping(id string) (error, response.Shipping)
	UpdateShipping(id string, input request.Shipping) (error, response.Shipping)
	DeleteShipping(id string) (error, response.Shipping)
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

func (ar *shippingRepository) GetAllShipping() (error, []response.Shipping) {
	var allShipping []models.Shipping
	var resAllShipping []response.Shipping
	err := ar.db.Find(&allShipping).Error
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}

	for i := 0; i < len(allShipping); i++ {
		shippingVm := domain.ConvertFromModelToShippingRes(allShipping[i])
		resAllShipping = append(resAllShipping, *shippingVm)
	}
	return nil, resAllShipping
}

func (ar *shippingRepository) GetShipping(id string) (error, response.Shipping) {
	var shippingData models.Shipping
	err := ar.db.First(&shippingData, "id = ?", id).Error

	if err != nil {
		return err, response.Shipping{}
	}

	return nil, *domain.ConvertFromModelToShippingRes(shippingData)
}

func (ar *shippingRepository) UpdateShipping(id string, input request.Shipping) (error, response.Shipping) {
	shippingData := models.Shipping{}
	err := ar.db.First(&shippingData, "id = ?", id).Error

	if err != nil {
		return err, response.Shipping{}
	}

	if input.Name != "" {
		shippingData.Name = input.Name
	}

	if err = ar.db.Save(&shippingData).Error; err != nil {
		return err, response.Shipping{}
	}
	return nil, *domain.ConvertFromModelToShippingRes(shippingData)
}

func (ar *shippingRepository) DeleteShipping(id string) (error, response.Shipping) {
	shippingData := models.Shipping{}
	res := response.Shipping{}
	find := ar.db.First(&shippingData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToShippingRes(shippingData)
	}
	err := ar.db.Delete(&shippingData, "id = ?", id).Error
	if err != nil {
		return err, response.Shipping{}
	}
	fmt.Printf("2")
	return nil, res
}
