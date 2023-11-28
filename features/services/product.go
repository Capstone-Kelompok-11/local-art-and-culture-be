package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
	"math"
)

type IProductService interface {
	CreateProduct(data *request.Product) (response.Product, error)
	GetAllProduct(nameFilter string, page, pageSize int) ([]response.Product, int, error)
	GetProduct(id string) (response.Product, error)
	UpdateProduct(id string, input request.Product) (response.Product, error)
	DeleteProduct(id string) (response.Product, error)
	CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	GetNextPage(currentPage, allPages int) int
	GetPrevPage(currentPage int) int
}

type ProductService struct {
	productRepository repositories.IProductRepository
}

func NewProductService(repo repositories.IProductRepository) *ProductService {
	return &ProductService{productRepository: repo}
}

func (pr *ProductService) CreateProduct(data *request.Product) (response.Product, error) {
	if data.Name == "" {
		return response.Product{}, errors.ERR_NAME_IS_EMPTY
	} 
	if data.Price == 0 {
		return response.Product{}, errors.ERR_PRICE_IS_EMPTY
	} 
	if data.Description == "" {
		return response.Product{}, errors.ERR_DESCRIPTION_IS_EMPTY
	} 
	if data.Status == "" {
		return response.Product{}, errors.ERR_STATUS_IS_EMPTY
	}

	res, err := pr.productRepository.CreateProduct(data)
	if err != nil {
		return response.Product{}, errors.ERR_CREATE_PRODUCT_DATABASE
	}
	return res, nil
}

func (ps *ProductService) GetAllProduct(nameFilter string, page, pageSize int) ([]response.Product, int, error) {
    res, allItems, err := ps.productRepository.GetAllProduct(nameFilter, page, pageSize)
    if err != nil {
        return nil, 0, err
    }
    return res, allItems, nil
}


func (pr *ProductService) GetProduct(id string) (response.Product, error) {
	if id == "" {
		return response.Product{}, errors.ERR_GET_PRODUCT_BAD_REQUEST_ID
	}
	res, err := pr.productRepository.GetProduct(id)
	if err != nil {
		return response.Product{}, err
	}
	return res, nil
}

func (pr *ProductService) UpdateProduct(id string, input request.Product) (response.Product, error) {
	if id == "" {
		return response.Product{}, errors.ERR_GET_PRODUCT_BAD_REQUEST_ID
	}
	res, err := pr.productRepository.UpdateProduct(id, input)
	if err != nil {
		return response.Product{}, err
	}
	return res, nil
}

func (pr *ProductService) DeleteProduct(id string) (response.Product, error) {
	if id == "" {
		return response.Product{}, errors.ERR_GET_PRODUCT_BAD_REQUEST_ID
	}
	res, err := pr.productRepository.DeleteProduct(id)
	if err != nil {
		return response.Product{}, err
	}
	return res, nil
}

func (pr *ProductService) CalculatePaginationValues(page, pageSize, allItmes int) (int, int) {
	pageInt := page
	if pageInt <= 0 {
		pageInt = 1
	}

	allPages := int(math.Ceil(float64(allItmes) / float64(pageSize)))

	if pageInt > allPages {
		pageInt = allPages
	}

	return pageInt, allPages
}

func (pr *ProductService) GetNextPage(currentPage, allPages int) int {
	if currentPage < allPages {
		return currentPage + 1
	}
	return allPages
}

func (pr *ProductService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}