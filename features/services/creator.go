package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ICreatorService interface {
	CreateCreator(data *request.Creator) (error, response.Creator)
	GetAllCreator() (error, []response.Creator)
	GetCreator(id string) (error, response.UserResponse)
	// UpdateCreator(id string, data request.Creator) (error, response.Creator)
	// DeleteCreator(id string) (error, response.Creator)
}

type CreatorService struct {
	creatorRepository repositories.ICreatorRepository
}

func NewCreatorService(repo repositories.ICreatorRepository) *CreatorService {
	return &CreatorService{creatorRepository: repo}
}

func (rs *CreatorService) CreateCreator(data *request.Creator) (error, response.Creator) {
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

	err, res := rs.creatorRepository.CreateCreator(data)
	if err != nil {
		return errors.ERR_REGISTER_USER_DATABASE, response.Creator{}
	}

	return nil, res
}

func (rs *CreatorService) GetAllCreator() (error, []response.Creator) {
	err, res := rs.creatorRepository.GetAllCreator()
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}
	return nil, res
}

func (rs *CreatorService) GetCreator(id string) (error, response.UserResponse) {
	if id == "" {
		return errors.ERR_GET_CREATOR_BAD_REQUEST_ID, response.UserResponse{}
	}
	err, res := rs.creatorRepository.GetCreator(id)
	if err != nil {
		return err, response.UserResponse{}
	}
	return nil, res
}
