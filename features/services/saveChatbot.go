package services

import (
	"lokasani/entity/models"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ISaveService interface {
	SaveChatbot(data models.SaveChatbot) (response.SaveChatbot, error)
	GetAllChatbot(UserId uint) ([]*response.SaveChatbot, error)
}

type SaveService struct {
	saveRepository repositories.ISaveRepository
}

func NewSaveService(repo repositories.ISaveRepository) *SaveService {
	return &SaveService{saveRepository: repo}
}

func (sv *SaveService) SaveChatbot(data models.SaveChatbot) (response.SaveChatbot, error) {
	if data.Message == "" {
		return response.SaveChatbot{}, errors.ERR_MESSAGE_IS_EMPTY
	}
	if data.Response == "" {
		return response.SaveChatbot{}, errors.ERR_RESPONSE_IS_EMPTY
	}

	res, err := sv.saveRepository.SaveChatbot(data)
	if err != nil {
		return response.SaveChatbot{}, errors.ERR_CREATE_SAVE_DATABASE
	}
	return res, nil
}

func (sv *SaveService) GetAllChatbot(UserId uint) ([]*response.SaveChatbot, error) {
	if UserId == 0 {
		return []*response.SaveChatbot{}, errors.ERR_GET_SAVE_BAD_REQUEST_ID
	}

	res, err := sv.saveRepository.GetAllChatbot(UserId)
	if err != nil {
		return []*response.SaveChatbot{}, errors.ERR_GET_DATA
	}
	return res, nil
}