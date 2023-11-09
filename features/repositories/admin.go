package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	CreateAdmin(data *request.Admin) (error, response.Admin)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (ar *adminRepository) CreateAdmin(data *request.Admin) (error, response.Admin) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Create(&dataAdmin).Error
	if err != nil {
		return err, response.Admin{}
	}
	return nil, *domain.ConvertFromModelToAdminRes(*dataAdmin)
}
