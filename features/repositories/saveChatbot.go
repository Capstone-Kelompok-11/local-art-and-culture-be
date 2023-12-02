package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

type ISaveRepository interface {
	SaveChatbot(data models.SaveChatbot) (response.SaveChatbot, error)
	GetAllChatbot(UserId uint) ([]*response.SaveChatbot, error)
}

type saveRepository struct {
	db *gorm.DB
}

func NewSaveRepository(db *gorm.DB) *saveRepository {
	return &saveRepository{db}
}

func (sv *saveRepository) SaveChatbot(data models.SaveChatbot) (response.SaveChatbot, error) {
	dataSave := domain.ConvertFromSaveReqToModel(data.Message, data.Response)
	err := sv.db.Create(&dataSave).Error
	if err != nil {
		return response.SaveChatbot{}, err
	}

	return *domain.ConvertFromModelToSaveRes(*dataSave), nil
}

func (sv *saveRepository) GetAllChatbot(UserId uint) ([]*response.SaveChatbot, error) {
	var saveData []models.SaveChatbot


	err := sv.db.Find(&saveData, "user_id = ?", UserId).Error
	if err != nil {
		return nil, err
	}

	resAllChat := domain.ListConvertFromModelToSaveRes(saveData)

	return resAllChat, nil
}
