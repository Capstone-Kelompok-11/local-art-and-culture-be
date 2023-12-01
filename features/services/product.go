package services

import (
	"encoding/csv"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
	"strconv"

	"strings"
)

type IProductService interface {
	CreateProduct(data *request.Product) ([]response.Product, error)
	GetAllProduct(nameFilter string) ([]response.Product, error)
	GetProduct(id string) (response.Product, error)
	UpdateProduct(id string, input request.Product) (response.Product, error)
	DeleteProduct(id string) (response.Product, error)
	getCategoryId(category string) uint
}

type ProductService struct {
	productRepository  repositories.IProductRepository
	categoryRepository repositories.ICategoryRepository
}

func NewProductService(productRepo repositories.IProductRepository, categoryRepo repositories.ICategoryRepository) *ProductService {
	return &ProductService{productRepository: productRepo, categoryRepository: categoryRepo}
}

func (pr *ProductService) CreateProduct(data *request.Product) ([]response.Product, error) {
	var result []response.Product
	if data.File != nil {
		csvData := string(data.File)

		reader := csv.NewReader(strings.NewReader(csvData))

		lines, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}
		var product request.Product
		for i := 1; i < len(lines); i++ {
			resultConvert := strings.Split(lines[i][0], ";")
			product.Name = resultConvert[0]
			product.Price, _ = strconv.ParseFloat(resultConvert[1], 64)
			product.Description = resultConvert[2]
			product.Status = resultConvert[3]
			product.CreatorId = 3
			catRes, err := pr.categoryRepository.GetAllCategory(resultConvert[4])
			if err != nil {
				return nil, errors.ERR_CATEGORY_NOT_FOUND
			}
			product.CategoryId = catRes[0].Id
			validation(product)
			res, err := pr.productRepository.CreateProduct(&product)
			result = append(result, res)
		}
	} else {
		err := validation(*data)
		if err != nil {
			return nil, err
		}
		res, err := pr.productRepository.CreateProduct(data)
		if err != nil {
			return nil, errors.ERR_CREATE_PRODUCT_DATABASE
		}
		result = append(result, res)
	}

	return result, nil
}

func (pr *ProductService) getCategoryId(category string) uint {
	res, err := pr.categoryRepository.GetAllCategory(category)
	if err != nil {
		return 0
	}
	return res[0].Id
}

func validation(data request.Product) error {
	if data.Name == "" {
		return errors.ERR_NAME_IS_EMPTY
	}
	if data.Price == 0 {
		return errors.ERR_PRICE_IS_EMPTY
	}
	if data.Description == "" {
		return errors.ERR_DESCRIPTION_IS_EMPTY
	}
	if data.Status == "" {
		return errors.ERR_STATUS_IS_EMPTY
	}
	return nil
}

func (pr *ProductService) GetAllProduct(nameFilter string) ([]response.Product, error) {
	res, err := pr.productRepository.GetAllProduct(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
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
