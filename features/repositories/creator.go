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
	GetAllCreator() (error, []response.Creator)
	GetCreator(id string) (error, response.UserResponse)
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

func (cr *creatorRepository) GetAllCreator() (error, []response.Creator) {
	var allCreator []models.Creator
	var resAllCreator []response.Creator
	err := cr.db.Preload("Creator").Find(&allCreator).Error
	if err != nil {
		return err, nil
	}

	for i := 0; i < len(allCreator); i++ {
		creatorVm := domain.ConvertFromModelToCreatorRes(allCreator[i])
		resAllCreator = append(resAllCreator, *creatorVm)
	}
	return nil, resAllCreator
}

func (cr *creatorRepository) GetCreator(id string) (error, response.UserResponse) {
	// var creatorData models.Creator
	var userData models.Users
	err := cr.db.Preload("Creator", "id = ?", id).First(&userData).Error

	if err != nil {
		return err, response.UserResponse{}
	}

	return nil, *domain.ConvertFromModelToUserRes(userData)
}
