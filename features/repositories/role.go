package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IRoleRepository interface {
	CreateRole(data *request.Role) (response.Role, error)
	GetAllRole(nameFilter string) ([]response.Role, error)
	GetRole(id string) (response.Role, error)
	UpdateRole(id string, input request.Role) (response.Role, error)
	DeleteRole(id string) (response.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (rr *roleRepository) CreateRole(data *request.Role) (response.Role, error) {
	dataRole := domain.ConvertFromRoleReqToModel(*data)
	err := rr.db.Create(&dataRole).Error
	if err != nil {
		return response.Role{}, err
	}
	return *domain.ConvertFromModelToRoleRes(*dataRole), nil
}

func (rr *roleRepository) GetAllRole(nameFilter string) ([]response.Role, error) {
	var allRole []models.Role
	var resAllRole []response.Role

	query := rr.db.Model(&models.Role{})
	if nameFilter != "" {
		query = query.Where("role LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allRole).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allRole); i++ {
		roleVm := domain.ConvertFromModelToRoleRes(allRole[i])
		resAllRole = append(resAllRole, *roleVm)
	}
	return resAllRole, nil
}

func (rr *roleRepository) GetRole(id string) (response.Role, error) {
	var roleData models.Role
	err := rr.db.First(&roleData, "id = ?", id).Error

	if err != nil {
		return response.Role{}, err
	}

	return *domain.ConvertFromModelToRoleRes(roleData), nil
}

func (rr *roleRepository) UpdateRole(id string, input request.Role) (response.Role, error) {
	roleData := models.Role{}
	err := rr.db.First(&roleData, "id = ?", id).Error

	if err != nil {
		return response.Role{}, errors.ERR_GET_ROLE_BAD_REQUEST_ID
	}

	if input.Role != "" {
		roleData.Role = input.Role
	}

	if err = rr.db.Save(&roleData).Error; err != nil {
		return response.Role{}, errors.ERR_UPDATE_DATA
	}
	return *domain.ConvertFromModelToRoleRes(roleData), nil
}

func (rr *roleRepository) DeleteRole(id string) (response.Role, error) {
	roleData := models.Role{}
	res := response.Role{}
	find := rr.db.First(&roleData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToRoleRes(roleData)
	}
	err := rr.db.Delete(&roleData, "id = ?", id).Error
	if err != nil {
		return response.Role{}, err
	}
	return res, nil
}