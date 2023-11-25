package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IRoleService interface {
	CreateRole(data *request.Role) (response.Role, error)
	GetAllRole(nameFilter string) ([]response.Role, error)
	GetRole(id string) (response.Role, error)
	UpdateRole(id string, data request.Role) (response.Role, error)
	DeleteRole(id string) (response.Role, error)
}

type RoleService struct {
	roleRepository repositories.IRoleRepository
}

func NewRoleService(repo repositories.IRoleRepository) *RoleService {
	return &RoleService{roleRepository: repo}
}

func (rs *RoleService) CreateRole(data *request.Role) (response.Role, error) {
	if data.Role == "" {
		return response.Role{}, errors.ERR_ROLE_IS_EMPTY
	}

	res, err := rs.roleRepository.CreateRole(data)
	if err != nil {
		return response.Role{}, errors.ERR_CREATE_ROLE
	}

	return res, nil
}

func (rs *RoleService) GetAllRole(nameFilter string) ([]response.Role, error) {
	res, err := rs.roleRepository.GetAllRole(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (rs *RoleService) GetRole(id string) (response.Role, error) {
	if id == "" {
		return response.Role{}, errors.ERR_GET_ROLE_BAD_REQUEST_ID
	}
	res, err := rs.roleRepository.GetRole(id)
	if err != nil {
		return response.Role{}, errors.ERR_GET_DATA
	}
	return res, nil
}

func (rs *RoleService) UpdateRole(id string, data request.Role) (response.Role, error) {
	if id == "" {
		return response.Role{}, errors.ERR_GET_ROLE_BAD_REQUEST_ID
	}
	res, err := rs.roleRepository.UpdateRole(id, data)
	if err != nil {
		return response.Role{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (rs *RoleService) DeleteRole(id string) (response.Role, error) {
	if id == "" {
		return response.Role{}, errors.ERR_GET_ROLE_BAD_REQUEST_ID
	}
	res, err := rs.roleRepository.DeleteRole(id)

	if err != nil {
		return response.Role{}, errors.ERR_DELETE_ROLE
	}

	return res, nil
}