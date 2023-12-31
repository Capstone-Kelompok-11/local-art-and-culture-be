package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

func ConvertFromAdminReqToModel(data request.SuperAdmin) *models.SuperAdmin {
	return &models.SuperAdmin{
		Model: gorm.Model{
			ID: data.Id,
		},
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		RoleId:		 data.RoleId,
	}
}

func ConvertFromModelToAdminRes(data models.SuperAdmin) *response.SuperAdmin {
	return &response.SuperAdmin{
		Id:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		RoleId:		 data.RoleId,
		Role:		 *ConvertFromModelToRoleRes(data.Role),
	}
}