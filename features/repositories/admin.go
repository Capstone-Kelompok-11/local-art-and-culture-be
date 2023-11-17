package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	RegisterAdmin(data *request.Admin) (response.Admin, error)
	LoginAdmin(data *request.Admin) (response.Admin, error)
	GetAllAdmin() ([]response.Admin, error)
	GetAdmin(id string) (response.Admin, error)
	UpdateAdmin(id string, input request.Admin) (response.Admin, error)
	DeleteAdmin(id string) (response.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (ar *adminRepository) RegisterAdmin(data *request.Admin) (response.Admin, error) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Create(&dataAdmin).Error
	if err != nil {
		return response.Admin{}, err
	}
	return *domain.ConvertFromModelToAdminRes(*dataAdmin), nil
}

func (ar *adminRepository) LoginAdmin(data *request.Admin) (response.Admin, error) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Where("email = ? ", dataAdmin.Email).First(&dataAdmin).Error
	if err != nil {
		return response.Admin{}, errors.ERR_EMAIL_NOT_FOUND
	}

	err = bcrypt.CheckPassword(data.Password, dataAdmin.Password)
	if err != nil {
		return response.Admin{}, errors.ERR_WRONG_PASSWORD
	}
	return *domain.ConvertFromModelToAdminRes(*dataAdmin), nil
}

func (ar *adminRepository) GetAllAdmin() ([]response.Admin, error) {
	var allAdmin []models.SuperAdmin
	var resAllAdmin []response.Admin
	err := ar.db.Find(&allAdmin).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allAdmin); i++ {
		adminVm := domain.ConvertFromModelToAdminRes(allAdmin[i])
		resAllAdmin = append(resAllAdmin, *adminVm)
	}
	return resAllAdmin, nil
}

func (ar *adminRepository) GetAdmin(id string) (response.Admin, error) {
	var adminData models.SuperAdmin
	err := ar.db.First(&adminData, "id = ?", id).Error

	if err != nil {
		return response.Admin{}, err
	}

	return *domain.ConvertFromModelToAdminRes(adminData), nil
}

func (ar *adminRepository) UpdateAdmin(id string, input request.Admin) (response.Admin, error) {
	adminData := models.SuperAdmin{}
	err := ar.db.First(&adminData, "id = ?", id).Error

	if err != nil {
		return response.Admin{}, err
	}

	if input.Email != "" {
		adminData.Email = input.Email
	}

	if input.Name != "" {
		adminData.Name = input.Name
	}

	if input.PhoneNumber != "" {
		adminData.PhoneNumber = input.PhoneNumber
	}

	if input.Password != "" {
		adminData.Password = input.Password
	}
	if err = ar.db.Save(&adminData).Error; err != nil {
		return response.Admin{}, err
	}
	return *domain.ConvertFromModelToAdminRes(adminData), nil
}

func (ar *adminRepository) DeleteAdmin(id string) (response.Admin, error) {
	adminData := models.SuperAdmin{}
	res := response.Admin{}
	find := ar.db.First(&adminData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToAdminRes(adminData)
	}
	err := ar.db.Delete(&adminData, "id = ?", id).Error
	if err != nil {
		return response.Admin{}, err
	}

	return res, nil
}
