package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IRoleService interface {
	CreateRole(data *request.Role) (error, response.Role)
	GetAllRole() (error, []response.Role)
	GetRole(id string) (error, response.Role)
	UpdateRole(id string, data request.Role) (error, response.Role)
	DeleteRole(id string) (error, response.Role)
}

type RoleService struct {
	roleRepository repositories.IRoleRepository
}

func NewRoleService(repo repositories.IRoleRepository) *RoleService {
	return &RoleService{roleRepository: repo}
}

func (rs *RoleService) CreateRole(data *request.Role) (error, response.Role) {
	if data.Role == "" {
		return errors.ERR_ROLE_IS_EMPTY, response.Role{}
	}

	err, res := rs.roleRepository.CreateRole(data)
	if err != nil {
		return errors.ERR_CREATE_ROLE, response.Role{}
	}

	return nil, res
}

func (rs *RoleService) GetAllRole() (error, []response.Role) {
	err, res := rs.roleRepository.GetAllRole()
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}
	return nil, res
}

func (rs *RoleService) GetRole(id string) (error, response.Role) {
	if id == "" {
		return errors.ERR_GET_ROLE_BAD_REQUEST_ID, response.Role{}
	}
	err, res := rs.roleRepository.GetRole(id)
	if err != nil {
		return errors.ERR_GET_DATA, response.Role{}
	}
	return nil, res
}

func (rs *RoleService) UpdateRole(id string, data request.Role) (error, response.Role) {
	if id == "" {
		return errors.ERR_GET_ROLE_BAD_REQUEST_ID, response.Role{}
	}
	err, res := rs.roleRepository.UpdateRole(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.Role{}
	}
	return nil, res
}

func (rs *RoleService) DeleteRole(id string) (error, response.Role) {
	if id == "" {
		return errors.ERR_GET_ROLE_BAD_REQUEST_ID, response.Role{}
	}
	err, res := rs.roleRepository.DeleteRole(id)

	if err != nil {
		return errors.ERR_DELETE_ROLE, response.Role{}
	}

	return nil, res
}
