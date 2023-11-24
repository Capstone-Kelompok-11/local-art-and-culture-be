package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromUserReqToModel(data request.User) *models.Users {
	return &models.Users{
		Model: gorm.Model{
			ID: data.Id,
		},
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		NIK: 		 data.NIK,
		Gender: 	 data.Gender,
		BirthDate:   data.BirthDate,
	}
}

func ConvertFromModelToUserRes(data models.Users) *response.User {
	return &response.User{
		Id:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		NIK: 		 data.NIK,
		Gender: 	 data.Gender,
		BirthDate:   data.BirthDate,
	}
}
