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
		return errors.ERR_GET_DATA, nil
	}

	for i := 0; i < len(allAdmin); i++ {
		adminVm := domain.ConvertFromModelToAdminRes(allAdmin[i])
		resAllAdmin = append(resAllAdmin, *adminVm)
	}
	return nil, resAllAdmin
}
