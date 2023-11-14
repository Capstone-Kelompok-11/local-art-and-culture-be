package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromUserReqToModel(data request.UserRequest) *models.Users {
	return &models.Users{
		Model: gorm.Model{
			ID: data.Id,
		},
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		BirthDate:   data.BirthDate,
	}
}

func ConvertFromModelToUserRes(data models.Users) *response.UserResponse {
	return &response.UserResponse{
		Id:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		BirthDate:   data.BirthDate,
	}
}


