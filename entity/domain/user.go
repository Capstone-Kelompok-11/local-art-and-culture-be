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
	return &response.User{
		Id:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Username: 	 data.Username,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		NIK: 		 data.NIK,
		Gender: 	 data.Gender,
		BirthDate:   data.BirthDate,
		RoleId: 	 data.RoleId,
		Role: 		 *ConvertFromModelToRoleRes(data.Role),
	}
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

// func ConvertFromModelToUserCRes(data models.Users) *response.UserCreatorResponse {
// 	return &response.UserCreatorResponse{
// 		Id:          data.ID,
// 		FirstName:   data.FirstName,
// 		LastName:    data.LastName,
// 		Username: 	 data.Username,
// 		Email:       data.Email,
// 		PhoneNumber: data.PhoneNumber,
// 		NIK: 		 data.NIK,
// 		Gender: 	 data.Gender,
// 		BirthDate:   data.BirthDate,
// 		RoleId: 	 data.RoleId,
// 		Role: 		 *ConvertFromModelToRoleRes(data.Role),
// 		Creator: 	*ConvertFromModelToCreatorRes(data.Cre),
// 	}
// }