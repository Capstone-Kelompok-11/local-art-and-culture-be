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

func ConvertFromModelToCreatorsRes(data models.Creator) *response.Creators {
	return &response.Creators{
		Id:          data.ID,
		Email:       data.Email,
		OutletName:  data.OutletName,
		PhoneNumber: data.PhoneNumber,
		UserId:      data.UserId,
		RoleId:      data.RoleId,
		Users:       *ConvertFromModelToUserRes(data.Users),
		Roles:       *ConvertFromModelToRoleRes(data.Roles),
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

func ConvertFromModelToUserCreatorRes(data models.Creator) *response.UserCreatorResponse {
	return &response.UserCreatorResponse{
		Id:          data.ID,
		FirstName:   data.Users.FirstName,
		LastName:    data.Users.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		BirthDate:   data.Users.BirthDate,
		Creator:     *ConvertFromModelToCreatorRes(data),
		// Creator:     *ConvertFromModelToCreatorRes(data.Creator),
	}
}