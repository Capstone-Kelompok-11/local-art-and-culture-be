package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
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

	res, err := rh.roleService.CreateRole(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) GetAllRole(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	res, err := rh.roleService.GetAllRole(nameFilter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) GetRole(c echo.Context) error {
	id := c.Param("id")
	res, err := rh.roleService.GetRole(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) UpdateRole(c echo.Context) error {
	id := c.Param("id")
	var input request.Role
	c.Bind(&input)
	res, err := rh.roleService.UpdateRole(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (rh *RoleHandler) DeleteRole(c echo.Context) error {
	id := c.Param("id")
	res, err := rh.roleService.DeleteRole(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}