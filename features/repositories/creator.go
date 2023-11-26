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
	GetAllCreator() ([]response.Creator, error)
	GetCreator(id string) (response.Creator, error)
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

func (cr *creatorRepository) GetAllCreator() ([]response.Creator, error) {
	var allCreator []models.Creator
	var resAllCreator []response.Creator
	err := cr.db.Preload("Users").Preload("Roles").Find(&allCreator).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allCreator); i++ {
		creatorVm := domain.ConvertFromModelToCreatorRes(allCreator[i])
		resAllCreator = append(resAllCreator, *creatorVm)
	}
	return resAllCreator, nil
}

func (cr *creatorRepository) GetCreator(id string) (response.Creator, error) {
	var userData models.Creator
	err := cr.db.Preload("Users").Preload("Roles").First(&userData, "id = ?", id).Error

	if err != nil {
		return response.Creator{}, err
	}

	return *domain.ConvertFromModelToCreatorRes(userData), nil
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
