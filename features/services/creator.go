package services

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ICreatorService interface {
	CreateCreator(data *request.Creator) (response.Creator, error)
	GetAllCreator(filter request.Creator) ([]response.Creators, error)
	GetAllCreatorByRole(filter request.Creator) ([]response.Creators, error)
	GetCreator(id string) (response.Creators, error)
	UpdateCreator(id string, data request.Creator) (response.Creator, error)
	DeleteCreator(id string) (response.Creator, error)
}

type CreatorService struct {
	creatorRepository repositories.ICreatorRepository
	roleRepo          repositories.IRoleRepository
}

func NewCreatorService(repo repositories.ICreatorRepository, role repositories.IRoleRepository) *CreatorService {
	return &CreatorService{
		creatorRepository: repo,
		roleRepo:          role,
	}
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
	if data.Roles == "" && data.RoleId == 0 {
		return response.Creator{}, errors.ERR_ROLE_IS_EMPTY
	}

	resRole, err := cs.roleRepo.GetAllRole(data.Roles)
	if err != nil {
		return response.Creator{}, errors.ERR_ROLE_IS_EMPTY
	}

	if len(resRole) == 0 {
		return response.Creator{}, errors.ERR_ROLE_IS_EMPTY
	}

	data.RoleId = resRole[0].Id
	fmt.Println(data.RoleId)
	fmt.Println(data)
	res, err := cs.creatorRepository.CreateCreator(data)
	if err != nil {
		return response.Creator{}, errors.ERR_CREATE_CREATOR_DATABASE
	}
	return res, nil
}

func (cs *CreatorService) GetAllCreator(filter request.Creator) ([]response.Creators, error) {
	res, err := cs.creatorRepository.GetAllCreator(filter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (cs *CreatorService) GetAllCreatorByRole(filter request.Creator) ([]response.Creators, error) {
	res, err := cs.creatorRepository.GetAllCreator(filter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (cs *CreatorService) GetCreator(id string) (response.Creators, error) {
	if id == "" {
		return response.Creators{}, errors.ERR_GET_CREATOR_BAD_REQUEST_ID
	}
	res, err := cs.creatorRepository.GetCreator(id)
	if err != nil {
		return response.Creators{}, err
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
