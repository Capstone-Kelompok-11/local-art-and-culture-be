package domain

import (
	"log"
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
		Username: 	 data.Username,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		NIK: 		 data.NIK,
		Gender: 	 data.Gender,
		BirthDate:   data.BirthDate,
		RoleId: 	 data.RoleId,
	}
}

func ConvertFromModelToUserRes(data models.Users) *response.User {
	userRes := &response.User{
		Id:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Username:    data.Username,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		NIK:         data.NIK,
		Gender:      data.Gender,
		BirthDate:   data.BirthDate,
		Role:        *ConvertFromModelToRoleRes(data.Role),
	}

	userRes.Date = data.CreatedAt.Format("2006-01-02 15:04:05")

	log.Printf("Before panic check: %v", data)
	if data.DeletedAt != nil && !data.DeletedAt.IsZero() {
    	log.Printf("Setting status: inactive")
    	userRes.Status = "inactive"
    	userRes.DeletedAt = *data.DeletedAt
	} else {
    	log.Printf("Setting status: active")
    	userRes.Status = "active"
	}	
	return userRes
}

func ConvertFromModelToUsersRes(data models.Users) *response.Users {
	return &response.Users{
		Id:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Username: 	 data.Username,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		NIK: 		 data.NIK,
		Gender: 	 data.Gender,
		BirthDate:   data.BirthDate,
	}
}

func UserCreateRequestToUserDomain(request request.User) *models.Users {
	return &models.Users{
		FirstName: 		request.FirstName,
		LastName:  		request.LastName,
		Username: 	 	request.Username,
		Email:     		request.Email,
		PhoneNumber:   	request.PhoneNumber,
		NIK:           	request.NIK,
		Gender:        	request.Gender,
		BirthDate:     	request.BirthDate,
		RoleId:        	request.RoleId, 
	}
}