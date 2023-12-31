package handler

import (
	"io"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	consts "lokasani/helpers/const"
	"lokasani/helpers/middleware"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService services.IProductService
}

func NewProductHandler(iProductService services.IProductService) *ProductHandler {
	return &ProductHandler{productService: iProductService}
}

func (pr *ProductHandler) CreateProduct(c echo.Context) error {
	_, role, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if role != consts.ProductCreator {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}
	// if roleId != 1 {
	// return response.NewErrorResponse(c, echo.ErrUnauthorized)
	// }

	var input request.Product
	c.Bind(&input)
	file, err := c.FormFile("file")
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return response.NewErrorResponse(c, err)
		}
		defer src.Close()

		fileBytes, err := io.ReadAll(src)
		input.File = fileBytes
	}
	res, err := pr.productService.CreateProduct(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *ProductHandler) GetAllProduct(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
    pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

    if page <= 0 {
        page = 1
    }

    if pageSize <= 0 {
        pageSize = 10
    }

	res, totalItems, err := pr.productService.GetAllProduct(nameFilter, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := pr.productService.CalculatePaginationValues(page, pageSize, totalItems)
	nextPage := pr.productService.GetNextPage(currentPage, allPages)
	prevPage := pr.productService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"allProduct": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":    allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
}

func (pr *ProductHandler) GetTrendingProduct(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
    pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

    if page <= 0 {
        page = 1
    }

    if pageSize <= 0 {
        pageSize = 10
    }

	res, totalItems, err := pr.productService.GetTrendingProduct(nameFilter, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := pr.productService.CalculatePaginationValues(page, pageSize, totalItems)
	nextPage := pr.productService.GetNextPage(currentPage, allPages)
	prevPage := pr.productService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"allProduct": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":    allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
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
	_, role, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if role != consts.ProductCreator {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}

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
	_, role, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if role != consts.ProductCreator {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}

	id := c.Param("id")
	res, err := pr.productService.DeleteProduct(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
