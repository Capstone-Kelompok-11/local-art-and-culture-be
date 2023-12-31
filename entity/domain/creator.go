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
		Address:     data.Address,
		// AddressId:   data.AddressId,
	}
}

func ConvertFromModelToCreatorsRes(data models.Creator) *response.Creators {
	return &response.Creators{
		Id:          data.ID,
		Email:       data.Email,
		OutletName:  data.OutletName,
		PhoneNumber: data.PhoneNumber,
		Roles: 		 data.Role.Role,
		Address: data.Address,
		Users: *ConvertFromModelToUsersRes(data.Users),
		Role:  *ConvertFromModelToRoleRes(data.Role),
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
		Address:     data.Address,
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
