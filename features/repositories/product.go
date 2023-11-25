package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(data *request.Product) (response.Product, error)
	GetAllProduct(nameFilter string) ([]response.Product, error)
	GetProduct(id string) (response.Product, error)
	UpdateProduct(id string, input request.Product) (response.Product, error)
	DeleteProduct(id string) (response.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (pr *productRepository) CreateProduct(data *request.Product) (response.Product, error) {
	dataProduct := domain.ConvertFromProductReqToModel(*data)
	err := pr.db.Create(&dataProduct).Error
	if err != nil {
		return response.Product{}, err
	}
	err = pr.db.Preload("Category").Preload("Creator").First(&dataProduct, "id = ?", dataProduct.ID).Error
	return *domain.ConvertFromModelToProductRes(*dataProduct), nil
}

func (pr *productRepository) GetAllProduct(nameFilter string) ([]response.Product, error) {
	var allProduct []models.Product
	var resAllProduct []response.Product

	query := pr.db.Preload("Category").Preload("Creator")

	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allProduct).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allProduct); i++ {
		productVm := domain.ConvertFromModelToProductRes(allProduct[i])
		resAllProduct = append(resAllProduct, *productVm)
	}

	return resAllProduct, nil
}

func (pr *productRepository) GetProduct(id string) (response.Product, error) {
	var productData models.Product
	err := pr.db.Preload("Category").Preload("Creator").First(&productData, "id = ?", id).Error

	if err != nil {
		return response.Product{}, err
	}
	return *domain.ConvertFromModelToProductRes(productData), nil
}

func (pr *productRepository) UpdateProduct(id string, input request.Product) (response.Product, error) {
	productData := models.Product{}
	err := pr.db.First(&productData, "id = ?", id).Error
	if err != nil {
		return response.Product{}, err
	}

	if input.Name != "" {
		productData.Name = input.Name
	}
	if input.Price != 0 {
		productData.Price = input.Price
	}
	if input.Description != "" {
		productData.Description = input.Description
	}
	if input.Status != "" {
		productData.Status = input.Status
	}

	if err = pr.db.Save(&productData).Error; err != nil {
		return response.Product{}, err
	}
	return *domain.ConvertFromModelToProductRes(productData), nil
}

func (pr *productRepository) DeleteProduct(id string) (response.Product, error) {
	productData := models.Product{}
	res := response.Product{}
	find := pr.db.Preload("Category").Preload("Creator").First(&productData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToProductRes(productData)
	}

	err := pr.db.Delete(&productData, "id = ?", id).Error
	if err != nil {
		return response.Product{}, err
	}
	return res, nil
}