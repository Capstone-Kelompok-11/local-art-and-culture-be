package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

type ICreatorRepository interface {
	CreateCreator(data *request.Creator) (error, response.Creator)
	GetAllCreator() (error, []response.UserCreatorResponse)
	GetCreator(id string) (error, response.UserCreatorResponse)
	// UpdateCreator(id string, input request.Creator) (error, response.Creator)
	// DeleteCreator(id string) (error, response.Creator)
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
