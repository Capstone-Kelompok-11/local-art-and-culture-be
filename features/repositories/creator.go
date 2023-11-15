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
	CreateCreator(data *request.Creator) (error, response.Creator)
	GetAllCreator() (error, []response.UserCreatorResponse)
	GetCreator(id string) (error, response.UserCreatorResponse)
	UpdateCreator(id string, input request.Creator) (error, response.Creator)
	DeleteCreator(id string) (error, response.Creator)
}

type creatorRepository struct {
	db *gorm.DB
}

func NewCreatorRepository(db *gorm.DB) *creatorRepository {
	return &creatorRepository{db}
}

func (cr *creatorRepository) CreateCreator(data *request.Creator) (error, response.Creator) {
	dataCreator := domain.ConvertFromCreatorReqToModel(*data)
	err := cr.db.Create(&dataCreator).Error
	if err != nil {
		return err, response.Creator{}
	}
	return nil, *domain.ConvertFromModelToCreatorRes(*dataCreator)
}

func (cr *creatorRepository) GetAllCreator() (error, []response.UserCreatorResponse) {
	var allCreator []models.Users
	var resAllCreator []response.UserCreatorResponse
	err := cr.db.Preload("Creator").Find(&allCreator).Error
	if err != nil {
		return err, nil
	}

	for i := 0; i < len(allCreator); i++ {
		creatorVm := domain.ConvertFromModelToUserCreatorRes(allCreator[i])
		resAllCreator = append(resAllCreator, *creatorVm)
	}
	return nil, resAllCreator
}

func (cr *creatorRepository) GetCreator(id string) (error, response.UserCreatorResponse) {
	var userData models.Users
	err := cr.db.Preload("Creator", "id = ?", id).First(&userData).Error

	if err != nil {
		return err, response.UserCreatorResponse{}
	}

	return nil, *domain.ConvertFromModelToUserCreatorRes(userData)
}

func (cr *creatorRepository) UpdateCreator(id string, input request.Creator) (error, response.Creator) {
	creatorData := models.Creator{}
	err := cr.db.First(&creatorData, "id = ?", id).Error

	if err != nil {
		return errors.ERR_GET_CREATOR_BAD_REQUEST_ID, response.Creator{}
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
		return errors.ERR_UPDATE_DATA, response.Creator{}
	}
	return nil, *domain.ConvertFromModelToCreatorRes(creatorData)
}

func (cr *creatorRepository) DeleteCreator(id string) (error, response.Creator) {
	creatorData := models.Creator{}
	res := response.Creator{}
	find := cr.db.First(&creatorData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToCreatorRes(creatorData)
	}
	err := cr.db.Delete(&creatorData, "id = ?", id).Error
	if err != nil {
		return err, response.Creator{}
	}

	return nil, res
}
