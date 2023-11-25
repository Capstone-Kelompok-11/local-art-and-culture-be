package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ICreatorRepository interface {
	CreateCreator(data *request.Creator) (response.Creator, error)
	GetAllCreator(nameFilter string) ([]response.UserCreatorResponse, error)
	GetCreator(id string) (response.UserCreatorResponse, error)
	UpdateCreator(id string, input request.Creator) (response.Creator, error)
	DeleteCreator(id string) (response.Creator, error)
}

type creatorRepository struct {
	db *gorm.DB
}

func NewCreatorRepository(db *gorm.DB) *creatorRepository {
	return &creatorRepository{db}
}

func (cr *creatorRepository) CreateCreator(data *request.Creator) (response.Creator, error) {
	dataCreator := domain.ConvertFromCreatorReqToModel(*data)
	err := cr.db.Create(&dataCreator).Error
	if err != nil {
		return response.Creator{}, err
	}
	err = cr.db.Preload("Users").Preload("Roles").First(&dataCreator, "id = ?", dataCreator.ID).Error
	return *domain.ConvertFromModelToCreatorRes(*dataCreator), nil
}

func (cr *creatorRepository) GetAllCreator(nameFilter string) ([]response.UserCreatorResponse, error) {
	var allCreator []models.Creator
	var resAllCreator []response.UserCreatorResponse

	query := cr.db.Preload("Roles").Preload("Users", func(db *gorm.DB) *gorm.DB {
		if nameFilter != "" {
			return db.Where("first_name LIKE ? OR last_name LIKE ?", "%"+nameFilter+"%", "%"+nameFilter+"%")
		}
		return db
	})

	err := query.Find(&allCreator).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allCreator); i++ {
		creatorVm := domain.ConvertFromModelToUserCreatorRes(allCreator[i])
		resAllCreator = append(resAllCreator, *creatorVm)
	}
	return resAllCreator, nil
}

func (cr *creatorRepository) GetCreator(id string) (response.UserCreatorResponse, error) {
	var userData models.Creator
	err := cr.db.Preload("Users").Preload("Roles").First(&userData, "id = ?", id).Error

	if err != nil {
		return response.UserCreatorResponse{}, err
	}

	return *domain.ConvertFromModelToUserCreatorRes(userData), nil
}

func (cr *creatorRepository) UpdateCreator(id string, input request.Creator) (response.Creator, error) {
	creatorData := models.Creator{}
	err := cr.db.Preload("Users").Preload("Roles").First(&creatorData, "id = ?", id).Error

	if err != nil {
		return response.Creator{}, errors.ERR_GET_CREATOR_BAD_REQUEST_ID
	}

	if input.OutletName != "" {
		creatorData.OutletName = input.OutletName
	}

	if input.Email != "" {
		creatorData.Email = input.Email
	}

	if input.PhoneNumber != "" {
		creatorData.PhoneNumber = input.PhoneNumber
	}

	if input.OutletName != "" {
		creatorData.OutletName = input.OutletName
	}

	if err = cr.db.Save(&creatorData).Error; err != nil {
		return response.Creator{}, errors.ERR_UPDATE_DATA
	}
	return *domain.ConvertFromModelToCreatorRes(creatorData), nil
}

func (cr *creatorRepository) DeleteCreator(id string) (response.Creator, error) {
	creatorData := models.Creator{}
	res := response.Creator{}
	find := cr.db.Preload("Users").Preload("Roles").First(&creatorData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToCreatorRes(creatorData)
	}
	err := cr.db.Delete(&creatorData, "id = ?", id).Error
	if err != nil {
		return response.Creator{}, err
	}

	return res, nil
}
