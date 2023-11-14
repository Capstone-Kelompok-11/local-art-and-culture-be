package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type RoleHandler struct {
	roleService services.IRoleService
}

func NewRoleHandler(iRoleService services.IRoleService) *RoleHandler {
	return &RoleHandler{roleService: iRoleService}
}

func (rh *RoleHandler) CreateRole(c echo.Context) error {
	var input request.Role
	c.Bind(&input)

	err, res := rh.roleService.CreateRole(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) GetAllRole(c echo.Context) error {
	err, res := rh.roleService.GetAllRole()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) GetRole(c echo.Context) error {
	id := c.Param("id")
	err, res := rh.roleService.GetRole(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) UpdateRole(c echo.Context) error {
	id := c.Param("id")
	var input request.Role
	c.Bind(&input)
	err, res := rh.roleService.UpdateRole(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) DeleteRole(c echo.Context) error {
	id := c.Param("id")
	err, res := rh.roleService.DeleteRole(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
