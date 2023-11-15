package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ICreatorService interface {
	CreateCreator(data *request.Creator) (error, response.Creator)
	GetAllCreator() (error, []response.UserCreatorResponse)
	GetCreator(id string) (error, response.UserCreatorResponse)
	UpdateCreator(id string, data request.Creator) (error, response.Creator)
	DeleteCreator(id string) (error, response.Creator)
}

type CreatorService struct {
	creatorRepository repositories.ICreatorRepository
}

func NewCreatorService(repo repositories.ICreatorRepository) *CreatorService {
	return &CreatorService{creatorRepository: repo}
}

func (cs *CreatorService) CreateCreator(data *request.Creator) (error, response.Creator) {
	if data.Email == "" {
		return errors.ERR_EMAIL_IS_EMPTY, response.Creator{}
	}
	if data.PhoneNumber == "" {
		return errors.ERR_PHONE_NUMBER_IS_EMPTY, response.Creator{}
	}
	if data.OutletName == "" {
		return errors.ERR_OUTLET_NAME_IS_EMPTY, response.Creator{}
	}
	if data.OutletName == "" {
		return errors.ERR_OUTLET_NAME_IS_EMPTY, response.Creator{}
	}

	err, res := cs.creatorRepository.CreateCreator(data)
	if err != nil {
		return errors.ERR_CREATE_CREATOR_DATABASE, response.Creator{}
	}

	return nil, res
}

func (cs *CreatorService) GetAllCreator() (error, []response.UserCreatorResponse) {
	err, res := cs.creatorRepository.GetAllCreator()
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}
	return nil, res
}

func (cs *CreatorService) GetCreator(id string) (error, response.UserCreatorResponse) {
	if id == "" {
		return errors.ERR_GET_CREATOR_BAD_REQUEST_ID, response.UserCreatorResponse{}
	}
	err, res := cs.creatorRepository.GetCreator(id)
	if err != nil {
		return err, response.UserCreatorResponse{}
	}
	return nil, res
}

func (cs *CreatorService) UpdateCreator(id string, data request.Creator) (error, response.Creator) {
	if id == "" {
		return errors.ERR_GET_CREATOR_BAD_REQUEST_ID, response.Creator{}
	}
	err, res := cs.creatorRepository.UpdateCreator(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.Creator{}
	}
	return nil, res
}

func (cs *CreatorService) DeleteCreator(id string) (error, response.Creator) {
	if id == "" {
		return errors.ERR_GET_CREATOR_BAD_REQUEST_ID, response.Creator{}
	}
	err, res := cs.creatorRepository.DeleteCreator(id)

	if err != nil {
		return errors.ERR_DELETE_CREATOR, response.Creator{}
	}

	return nil, res
}
