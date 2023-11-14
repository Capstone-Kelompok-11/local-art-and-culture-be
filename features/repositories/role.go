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
	CreateRole(data *request.Role) (error, response.Role)
	GetAllRole() (error, []response.Role)
	GetRole(id string) (error, response.Role)
	UpdateRole(id string, input request.Role) (error, response.Role)
	DeleteRole(id string) (error, response.Role)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (rr *roleRepository) CreateRole(data *request.Role) (error, response.Role) {
	dataRole := domain.ConvertFromRoleReqToModel(*data)
	err := rr.db.Create(&dataRole).Error
	if err != nil {
		return err, response.Role{}
	}
	return nil, *domain.ConvertFromModelToRoleRes(*dataRole)
}

func (rr *roleRepository) GetAllRole() (error, []response.Role) {
	var allRole []models.Role
	var resAllRole []response.Role
	err := rr.db.Find(&allRole).Error
	if err != nil {
		return err, nil
	}

	for i := 0; i < len(allRole); i++ {
		roleVm := domain.ConvertFromModelToRoleRes(allRole[i])
		resAllRole = append(resAllRole, *roleVm)
	}
	return nil, resAllRole
}

func (rr *roleRepository) GetRole(id string) (error, response.Role) {
	var roleData models.Role
	err := rr.db.First(&roleData, "id = ?", id).Error

	if err != nil {
		return err, response.Role{}
	}

	return nil, *domain.ConvertFromModelToRoleRes(roleData)
}

func (rr *roleRepository) UpdateRole(id string, input request.Role) (error, response.Role) {
	roleData := models.Role{}
	err := rr.db.First(&roleData, "id = ?", id).Error

	if err != nil {
		return errors.ERR_GET_ROLE_BAD_REQUEST_ID, response.Role{}
	}

	if input.Role != "" {
		roleData.Role = input.Role
	}

	if err = rr.db.Save(&roleData).Error; err != nil {
		return errors.ERR_UPDATE_DATA, response.Role{}
	}
	return nil, *domain.ConvertFromModelToRoleRes(roleData)
}

func (rr *roleRepository) DeleteRole(id string) (error, response.Role) {
	roleData := models.Role{}
	res := response.Role{}
	find := rr.db.First(&roleData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToRoleRes(roleData)
	}
	err := rr.db.Delete(&roleData, "id = ?", id).Error
	if err != nil {
		return err, response.Role{}
	}

	return nil, res
}
