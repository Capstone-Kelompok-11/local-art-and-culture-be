package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ICategoryService interface {
	CreateCategory(data *request.Category) (response.Category, error)
	GetAllCategory(nameFilter string) ([]response.Category, error)
	GetCategory(id string) (response.Category, error)
	GetTypeCategory(Type string) (response.Category, error)
	UpdateCategory(id string, input request.Category) (response.Category, error)
	DeleteCategory(id string) (response.Category, error)
}

type CategoryService struct {
	categoryRepository repositories.ICategoryRepository
}

func NewCategoryService(repo repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: repo}
}

func (ca *CategoryService) CreateCategory(data *request.Category) (response.Category, error) {
	if data.Category == "" {
		return response.Category{}, errors.ERR_CATEGORY_IS_EMPTY
	} else if data.Type == "" {
		return response.Category{}, errors.ERR_TYPE_IS_EMPTY
	}

	res, err := ca.categoryRepository.CreateCategory(data)
	if err != nil {
		return response.Category{}, errors.ERR_CREATE_CATEGORY_DATABASE
	}
	return res, nil
}

func (ca *CategoryService) GetAllCategory(nameFilter string) ([]response.Category, error) {
	res, err := ca.categoryRepository.GetAllCategory(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (ca *CategoryService) GetCategory(id string) (response.Category, error) {
	if id == "" {
		return response.Category{}, errors.ERR_GET_CATEGORY_BAD_REQUEST_ID
	}
	
	res, err := ca.categoryRepository.GetCategory(id)
	if err != nil {
		return response.Category{}, err
	}
	return res, nil
}

func (ca *CategoryService) GetTypeCategory(Type string) (response.Category, error) {
	if Type == "" {
		return response.Category{}, errors.ERR_GET_CATEGORY_BAD_REQUEST_TYPE
	}
	
	res, err := ca.categoryRepository.GetTypeCategory(Type)
	if err != nil {
		return response.Category{}, err
	}
	return res, nil
}

func (ca *CategoryService) UpdateCategory(id string, input request.Category) (response.Category, error) {
	if id == "" {
		return response.Category{}, errors.ERR_GET_CATEGORY_BAD_REQUEST_ID
	}

	res, err := ca.categoryRepository.UpdateCategory(id, input)
	if err != nil {
		return response.Category{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (ca *CategoryService) DeleteCategory(id string) (response.Category, error) {
	if id == "" {
		return response.Category{}, errors.ERR_GET_CATEGORY_BAD_REQUEST_ID
	}

	res, err := ca.categoryRepository.DeleteCategory(id)
	if err != nil {
		return response.Category{}, errors.ERR_DELETE_CATEGORY
	}
	return res, nil
}