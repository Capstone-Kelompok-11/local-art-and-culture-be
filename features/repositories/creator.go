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
	GetAllCreator() ([]response.UserCreatorResponse, error)
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
	return *domain.ConvertFromModelToCreatorRes(*dataCreator), nil
}

func (cr *creatorRepository) GetAllCreator() ([]response.UserCreatorResponse, error) {
	var allCreator []models.Users
	var resAllCreator []response.UserCreatorResponse
	err := cr.db.Preload("Creator").Find(&allCreator).Error
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
	var userData models.Users
	err := cr.db.Preload("Creator", "id = ?", id).First(&userData).Error

	if err != nil {
		return response.UserCreatorResponse{}, err
	}

	return *domain.ConvertFromModelToUserCreatorRes(userData), nil
}

func (cr *creatorRepository) UpdateCreator(id string, input request.Creator) (response.Creator, error) {
	creatorData := models.Creator{}
	err := cr.db.First(&creatorData, "id = ?", id).Error

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
	find := cr.db.First(&creatorData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToCreatorRes(creatorData)
	}
	err := cr.db.Delete(&creatorData, "id = ?", id).Error
	if err != nil {
		return response.Creator{}, err
	}

	return res, nil
}