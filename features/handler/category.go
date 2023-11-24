package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type CategoryHandler struct {
	categoryService services.ICategoryService
}

func NewCategoryHandler(iCategoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: &iCategoryService}
}

func (ca *CategoryHandler) CreateCategory(c echo.Context) error {
	var input request.Category
	c.Bind(&input)

	res, err := ca.categoryService.CreateCategory(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ca *CategoryHandler) GetAllCategory(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	res, err := ca.categoryService.GetAllCategory(nameFilter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ca *CategoryHandler) GetCategory(c echo.Context) error {
	id := c.Param("id")
	res, err := ca.categoryService.GetCategory(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ca *CategoryHandler) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	var input request.Category
	c.Bind(&input)

	res, err := ca.categoryService.UpdateCategory(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ca *CategoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	res, err := ca.categoryService.DeleteCategory(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}