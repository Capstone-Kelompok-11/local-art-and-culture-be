package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IAddressRepository interface {
	CreateAddress(data *request.Address) (response.Address, error)
	GetAllAddress(nameFilter string) ([]response.Address, error)
	GetAddress(id string) (response.Address, error)
	UpdateAddress(id string, input request.Address) (response.Address, error)
	DeleteAddress(id string) (response.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *addressRepository {
	return &addressRepository{db}
}

func (ad *addressRepository) CreateAddress(data *request.Address) (response.Address, error) {
	dataAddress := domain.ConvertFromAddressReqToModel(*data)
	err := ad.db.Create(&dataAddress).Error
	if err != nil {
		return response.Address{}, err
	}
	//err = ad.db.Preload("Address").First(&dataAddress, "id = ?", dataAddress.ID).Error
	return *domain.ConvertFromModelToAddressRes(*dataAddress), nil
}

func (ad *addressRepository) GetAllAddress(nameFilter string) ([]response.Address, error) {
	var allAddress []models.Address
	var resAllAddress []response.Address

	query := ad.db.Model(&models.Address{})
	if nameFilter != "" {
		query = query.Where("address LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allAddress).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allAddress); i++ {
		AddressVm := domain.ConvertFromModelToAddressRes(allAddress[i])
		resAllAddress = append(resAllAddress, *AddressVm)
	}
	return resAllAddress, nil
}

func (ad *addressRepository) GetAddress(id string) (response.Address, error) {
	var addressData models.Address
	err := ad.db.First(&addressData, "id = ?", id).Error

	if err != nil {
		return response.Address{}, err
	}

	return *domain.ConvertFromModelToAddressRes(addressData), nil
}

func (ad *addressRepository) UpdateAddress(id string, input request.Address) (response.Address, error) {
	AddressData := models.Address{}
	err := ad.db.First(&AddressData, "id = ?", id).Error

	if err != nil {
		return response.Address{}, errors.ERR_GET_BAD_REQUEST_ID
	}

	if input.Address != "" {
		AddressData.Address = input.Address
	}

	if err = ad.db.Save(&AddressData).Error; err != nil {
		return response.Address{}, errors.ERR_UPDATE_DATA
	}
	return *domain.ConvertFromModelToAddressRes(AddressData), nil
}

func (ad *addressRepository) DeleteAddress(id string) (response.Address, error) {
	addressData := models.Address{}
	res := response.Address{}
	find := ad.db.First(&addressData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToAddressRes(addressData)
	}
	err := ad.db.Delete(&addressData, "id = ?", id).Error
	if err != nil {
		return response.Address{}, err
	}
	return res, nil
}