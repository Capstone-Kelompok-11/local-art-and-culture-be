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
	RegisterAdmin(data *request.SuperAdmin) (response.SuperAdmin, error)
	LoginAdmin(data *request.SuperAdmin) (response.SuperAdmin, error)
	GetAllAdmin(nameFilter string, page, pageSize int) ([]response.SuperAdmin, int, error)
	GetAdmin(id string) (response.SuperAdmin, error)
	UpdateAdmin(id string, input request.SuperAdmin) (response.SuperAdmin, error)
	DeleteAdmin(id string) (response.SuperAdmin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (ar *adminRepository) RegisterAdmin(data *request.SuperAdmin) (response.SuperAdmin, error) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Create(&dataAdmin).Error
	if err != nil {
		return response.SuperAdmin{}, err
	}
	return *domain.ConvertFromModelToAdminRes(*dataAdmin), nil
}

func (ar *adminRepository) LoginAdmin(data *request.SuperAdmin) (response.SuperAdmin, error) {
	dataAdmin := domain.ConvertFromAdminReqToModel(*data)
	err := ar.db.Where("email = ? ", dataAdmin.Email).First(&dataAdmin).Error
	if err != nil {
		return response.SuperAdmin{}, errors.ERR_EMAIL_NOT_FOUND
	}

	err = bcrypt.CheckPassword(data.Password, dataAdmin.Password)
	if err != nil {
		return response.SuperAdmin{}, errors.ERR_WRONG_PASSWORD
	}
	return *domain.ConvertFromModelToAdminRes(*dataAdmin), nil
}

func (ar *adminRepository) GetAllAdmin(nameFilter string, page, pageSize int) ([]response.SuperAdmin, int, error){
	var allAdmin []models.SuperAdmin
	var resAllAdmin []response.SuperAdmin

	query := ar.db.Model(&models.SuperAdmin{})
	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	offset := (page - 1) * pageSize

	query = query.Limit(pageSize).Offset(offset)
	
	err := query.Find(&allAdmin).Error
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(allAdmin); i++ {
		adminVm := domain.ConvertFromModelToAdminRes(allAdmin[i])
		resAllAdmin = append(resAllAdmin, *adminVm)
	}

	var allItems int64
	query.Count(&allItems)

	return resAllAdmin, int(allItems), nil
}

func (ar *adminRepository) GetAdmin(id string) (response.SuperAdmin, error) {
	var adminData models.SuperAdmin
	err := ar.db.First(&adminData, "id = ?", id).Error

	if err != nil {
		return response.SuperAdmin{}, err
	}

	return *domain.ConvertFromModelToAdminRes(adminData), nil
}

func (ar *adminRepository) UpdateAdmin(id string, input request.SuperAdmin) (response.SuperAdmin, error) {
	adminData := models.SuperAdmin{}
	err := ar.db.First(&adminData, "id = ?", id).Error

	if err != nil {
		return response.SuperAdmin{}, err
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
		return response.SuperAdmin{}, err
	}
	return *domain.ConvertFromModelToAdminRes(adminData), nil
}

func (ar *adminRepository) DeleteAdmin(id string) (response.SuperAdmin, error) {
	adminData := models.SuperAdmin{}
	res := response.SuperAdmin{}
	find := ar.db.First(&adminData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToAdminRes(adminData)
	}
	err := ar.db.Delete(&adminData, "id = ?", id).Error
	if err != nil {
		return response.SuperAdmin{}, err
	}

	return res, nil
}