package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromCreatorReqToModel(data request.Creator) *models.Creator {
	return &models.Creator{
		Model: gorm.Model{
			ID: data.Id,
		},
		Email:       data.Email,
		OutletName:  data.OutletName,
		PhoneNumber: data.PhoneNumber,
		UserId:      data.UserId,
		RoleId:      data.RoleId,
		// AddressId:   data.AddressId,
	}
}

func ConvertFromModelToCreatorRes(data models.Creator) *response.Creator {
	return &response.Creator{
		Id:          data.ID,
		Email:       data.Email,
		OutletName:  data.OutletName,
		PhoneNumber: data.PhoneNumber,
		UserId:      data.UserId,
		RoleId:      data.RoleId,
		// AddressId:   data.AddressId,
	}
}

func ConvertFromModelToUserCreatorRes(data models.Users) *response.UserCreatorResponse {
	return &response.UserCreatorResponse{
		Id:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		BirthDate:   data.BirthDate,
		Creator:     *ConvertFromModelToCreatorRes(data.Creator),
	}
}
