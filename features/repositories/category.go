package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	CreateCategory(data *request.Category) (response.Category, error)
	GetAllCategory(nameFilter string) ([]response.Category, error)
	GetCategory(id string) (response.Category, error)
	UpdateCategory(id string, input request.Category) (response.Category, error)
	DeleteCategory(id string) (response.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (ca *categoryRepository) CreateCategory(data *request.Category) (response.Category, error) {
	dataCategory := domain.ConvertFromCategoryReqToModel(*data)
	err := ca.db.Create(&dataCategory).Error
	if err != nil {
		return response.Category{}, err
	}
	return *domain.ConvertFromModelToCategoryRes(*dataCategory), nil
}

func (ca *categoryRepository) GetAllCategory(nameFilter string) ([]response.Category, error) {
	var allCategory []models.Category
	var resAllCategory []response.Category

	query := ca.db.Model(&models.Category{})
	if nameFilter != "" {
		query = query.Where("category LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allCategory).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allCategory); i++ {
		categoryVm := domain.ConvertFromModelToCategoryRes(allCategory[i])
		resAllCategory = append(resAllCategory, *categoryVm)
	}
	return resAllCategory, nil
}

func (ca *categoryRepository) GetCategory(id string) (response.Category, error) {
	var categoryData models.Category
	err := ca.db.First(&categoryData, "id = ?", id).Error
	if err != nil {
		return response.Category{}, err
	}
	return *domain.ConvertFromModelToCategoryRes(categoryData), nil
}

func (ca *categoryRepository) UpdateCategory(id string, input request.Category) (response.Category, error) {
	categoryData := models.Category{}
	err := ca.db.First(&categoryData, "id = ?", id).Error
	if err != nil {
		return response.Category{}, err
	}

	if input.Category != "" {
		categoryData.Category = input.Category
	} else if input.Type != "" {
		categoryData.Type = input.Type
	}

	if err = ca.db.Save(&categoryData).Error; err != nil {
		return response.Category{}, err
	}
	return *domain.ConvertFromModelToCategoryRes(categoryData),  nil
}

func (ca *categoryRepository) DeleteCategory(id string) (response.Category, error) {
	categoryData := models.Category{}
	res := response.Category{} 
	find := ca.db.First(&categoryData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToCategoryRes(categoryData)
	}
	err := ca.db.Delete(&categoryData, "id = ?", id).Error
	if err != nil {
		return response.Category{}, err
	}
	return res, nil
}