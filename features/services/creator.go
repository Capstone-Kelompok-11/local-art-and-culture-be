package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ICreatorService interface {
	CreateCreator(data *request.Creator) (response.Creator, error)
	GetAllCreator(nameFilter string) ([]response.UserCreatorResponse, error)
	GetCreator(id string) (response.UserCreatorResponse, error)
	UpdateCreator(id string, data request.Creator) (response.Creator, error)
	DeleteCreator(id string) (response.Creator, error)
}

type CreatorService struct {
	creatorRepository repositories.ICreatorRepository
}

func NewCreatorService(repo repositories.ICreatorRepository) *CreatorService {
	return &CreatorService{creatorRepository: repo}
}

func (cs *CreatorService) CreateCreator(data *request.Creator) (response.Creator, error) {
	if data.Email == "" {
		return response.Creator{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.PhoneNumber == "" {
		return response.Creator{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}
	if data.OutletName == "" {
		return response.Creator{}, errors.ERR_OUTLET_NAME_IS_EMPTY
	}
	if data.OutletName == "" {
		return response.Creator{}, errors.ERR_OUTLET_NAME_IS_EMPTY
	}

	res, err := cs.creatorRepository.CreateCreator(data)
	if err != nil {
		return response.Creator{}, errors.ERR_CREATE_CREATOR_DATABASE
	}

	return res, nil
}

func (cs *CreatorService) GetAllCreator(nameFilter string) ([]response.UserCreatorResponse, error) {
	res, err := cs.creatorRepository.GetAllCreator(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (cs *CreatorService) GetCreator(id string) (response.UserCreatorResponse, error) {
	if id == "" {
		return response.UserCreatorResponse{}, errors.ERR_GET_CREATOR_BAD_REQUEST_ID
	}
	res, err := cs.creatorRepository.GetCreator(id)
	if err != nil {
		return response.UserCreatorResponse{}, err
	}
	return res, nil
}

func (cs *CreatorService) UpdateCreator(id string, data request.Creator) (response.Creator, error) {
	if id == "" {
		return response.Creator{}, errors.ERR_GET_CREATOR_BAD_REQUEST_ID
	}
	res, err := cs.creatorRepository.UpdateCreator(id, data)
	if err != nil {
		return response.Creator{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (cs *CreatorService) DeleteCreator(id string) (response.Creator, error) {
	if id == "" {
		return response.Creator{}, errors.ERR_GET_CREATOR_BAD_REQUEST_ID
	}
	res, err := cs.creatorRepository.DeleteCreator(id)

	if err != nil {
		return response.Creator{}, errors.ERR_DELETE_CREATOR
	}

	return res, nil
}
