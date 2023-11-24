package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	productService services.IProductService
}

func NewProductHandler(iProductService services.IProductService) *ProductHandler {
	return &ProductHandler{productService: iProductService}
}

func (pr *ProductHandler) CreateProduct(c echo.Context) error {
	var input request.Product
	c.Bind(&input)
	res, err := pr.productService.CreateProduct(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *ProductHandler) GetAllProduct(c echo.Context) error {
	nameFilter := c.QueryParam("name")

    res, err := pr.productService.GetAllProduct(nameFilter)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}

func (pr *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")
	res, err := pr.productService.GetProduct(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var input request.Product
	c.Bind(&input)

	res, err := pr.productService.UpdateProduct(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	res, err := pr.productService.DeleteProduct(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
