package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromRoleReqToModel(data request.Role) *models.Role {
	return &models.Role{
		Model: gorm.Model{
			ID: data.Id,
		},
		Role: data.Role,
	}
}

func ConvertFromModelToRoleRes(data models.Role) *response.Role {
	return &response.Role{
		Id:   data.ID,
		Role: data.Role,
	}
}
