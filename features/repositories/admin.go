package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	RegisterAdmin(data *request.Admin) (error, response.Admin)
	LoginAdmin(data *request.Admin) (error, response.Admin)	
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
