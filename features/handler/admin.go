package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	consts "lokasani/helpers/const"
	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService services.IAdminService
}

func NewAdminHandler(iAdminService services.IAdminService) *AdminHandler {
	return &AdminHandler{adminService: iAdminService}
}

func (ah *AdminHandler) RegisterAdmin(c echo.Context) error {
	var input request.SuperAdmin
	c.Bind(&input)

	res, err := ah.adminService.RegisterAdmin(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	// if res.role
	token, err := middleware.CreateToken(uint(res.Id), consts.AdminRole, 0)
	if err != nil {
		return response.NewErrorResponse(c, errors.ERR_TOKEN)
	}
	res.Token = token

	middleware.SetTokenCookie(c, token)

	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) LoginAdmin(c echo.Context) error {
	var input request.SuperAdmin
	c.Bind(&input)

	res, err := ah.adminService.LoginAdmin(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	token, err := middleware.CreateToken(0, consts.AdminRole, uint(res.Id))
    if err != nil {
        return response.NewErrorResponse(c, errors.ERR_TOKEN)
    }
    res.Token = token

	middleware.SetTokenCookie(c, token)
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) GetAllAdmin(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	page, pageSize := 1, 10

	res, allItems, err := ah.adminService.GetAllAdmin(nameFilter, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := ah.adminService.CalculatePaginationValues(page, pageSize, allItems)
	nextPage := ah.adminService.GetNextPage(currentPage, allPages)
	prevPage := ah.adminService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"data": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":  allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
}

func (ah *AdminHandler) GetAdmin(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.adminService.GetAdmin(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) UpdateAdmin(c echo.Context) error {
	id := c.Param("id")
	var input request.SuperAdmin
	c.Bind(&input)
	res, err := ah.adminService.UpdateAdmin(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) DeleteAdmin(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.adminService.DeleteAdmin(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}