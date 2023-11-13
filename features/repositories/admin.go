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
	RegisterAdmin(data *request.Admin) (error, response.Admin)
	LoginAdmin(data *request.Admin) (error, response.Admin)
	GetAllAdmin() (error, []response.Admin)
	GetAdmin(id string) (error, response.Admin)
	UpdateAdmin(id string, input request.Admin) (error, response.Admin)
	DeleteAdmin(id string) (error, response.Admin)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (ar *adminRepository) RegisterAdmin(data *request.Admin) (error, response.Admin) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Create(&dataAdmin).Error
	if err != nil {
		return err, response.Admin{}
	}
	return nil, *domain.ConvertFromModelToAdminRes(*dataAdmin)
}

func (ar *adminRepository) LoginAdmin(data *request.Admin) (error, response.Admin) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Where("email = ? ", dataAdmin.Email).First(&dataAdmin).Error
	if err != nil {
		return errors.ERR_EMAIL_NOT_FOUND, response.Admin{}
	}

	err = bcrypt.CheckPassword(data.Password, dataAdmin.Password)
	if err != nil {
		return errors.ERR_WRONG_PASSWORD, response.Admin{}
	}
	return nil, *domain.ConvertFromModelToAdminRes(*dataAdmin)
}

func (ar *adminRepository) GetAllAdmin() (error, []response.Admin) {
	var allAdmin []models.SuperAdmin
	var resAllAdmin []response.Admin
	err := ar.db.Find(&allAdmin).Error
	if err != nil {
		return err, nil
	}

	for i := 0; i < len(allAdmin); i++ {
		adminVm := domain.ConvertFromModelToAdminRes(allAdmin[i])
		resAllAdmin = append(resAllAdmin, *adminVm)
	}
	return nil, resAllAdmin
}

func (ar *adminRepository) GetAdmin(id string) (error, response.Admin) {
	var adminData models.SuperAdmin
	err := ar.db.First(&adminData, "id = ?", id).Error

	if err != nil {
		return err, response.Admin{}
	}

	return nil, *domain.ConvertFromModelToAdminRes(adminData)
}

func (ar *adminRepository) UpdateAdmin(id string, input request.Admin) (error, response.Admin) {
	adminData := models.SuperAdmin{}
	err := ar.db.First(&adminData, "id = ?", id).Error

	if err != nil {
		return err, response.Admin{}
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
		return err, response.Admin{}
	}
	return nil, *domain.ConvertFromModelToAdminRes(adminData)
}

func (ar *adminRepository) DeleteAdmin(id string) (error, response.Admin) {
	adminData := models.SuperAdmin{}
	res := response.Admin{}
	find := ar.db.First(&adminData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToAdminRes(adminData)
	}
	err := ar.db.Delete(&adminData, "id = ?", id).Error
	if err != nil {
		return err, response.Admin{}
	}

	return nil, res
}
