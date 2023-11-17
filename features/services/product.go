package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IProductService interface {
	CreateProduct(data *request.Product) (response.Product, error)
	GetAllProduct() ([]response.Product, error)
	GetProduct(id string) (response.Product, error)
	UpdateProduct(id string, input request.Product) (response.Product, error)
	DeleteProduct(id string) (response.Product, error)
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
	} else if data.Price == "" {
		return response.Product{}, errors.ERR_PRICE_IS_EMPTY
	} else if data.Description == "" {
		return response.Product{}, errors.ERR_DESCRIPTION_IS_EMPTY
	} else if data.Status == "" {
		return response.Product{}, errors.ERR_STATUS_IS_EMPTY
	}

	res, err := pr.productRepository.CreateProduct(data)
	if err != nil {
		return response.Product{}, errors.ERR_CREATE_PRODUCT_DATABASE
	}
	return res, nil
}

func (pr *ProductService) GetAllProduct() ([]response.Product, error) {
	res, err := pr.productRepository.GetAllProduct()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (pr *ProductService) GetProduct(id string) (response.Product, error) {
	if id == "" {
		return response.Product{}, errors.ERR_GET_PRODUCT_BAD_REQUEST_ID
	}
	res, err := pr.GetProduct(id)
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