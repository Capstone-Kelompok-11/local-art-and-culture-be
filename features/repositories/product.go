package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"
	"sort"
	"time"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(data *request.Product) (response.Product, error)
	GetAllProduct(nameFilter string, page, pageSize int) ([]response.Product, int, error)
	GetTrendingProduct(nameFilter string, page, pageSize int) ([]response.Products, int, error)
	GetProduct(id string) (response.Product, error)
	UpdateProduct(id string, input request.Product) (response.Product, error)
	DeleteProduct(id string) (response.Product, error)
	GetTotalLikes(SourceId uint, SourceStr string) (uint, error)
	GetTotalLikesLastTwoWeeks(SourceId uint, SourceStr string) (uint, error)
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

func (pr *productRepository) GetAllProduct(nameFilter string, page, pageSize int) ([]response.Product, int, error) {
	var allProduct []models.Product
	var resAllProduct []response.Product

	query := pr.db.Preload("Category").Preload("Creator")

	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	offset := (page - 1) * pageSize

	query = query.Limit(pageSize).Offset(offset)

	err := query.Find(&allProduct).Error
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allProduct); i++ {
		productVm := domain.ConvertFromModelToProductRes(allProduct[i])
		resAllProduct = append(resAllProduct, *productVm)
	}

	var totalItems int64
	query.Count(&totalItems)

	return resAllProduct, int(totalItems), nil
}

func (pr *productRepository) GetTrendingProduct(nameFilter string, page, pageSize int) ([]response.Products, int, error) {
	var allProduct []models.Product
	var resAllProduct []response.Products

	query := pr.db.Preload("Category").Preload("Creator").Preload("Like")

	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	offset := (page - 1) * pageSize

	query = query.Limit(pageSize).Offset(offset)

	twoWeeksAgo := time.Now().Add(-2 * 7 * 24 * time.Hour)
	err := query.Where("created_at >= ?", twoWeeksAgo).Find(&allProduct).Error
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allProduct); i++ {
		productVm := domain.ConvertFromModelToProductsRes(allProduct[i])

		totalLikes, err := pr.GetTotalLikes(allProduct[i].ID, "product")
		if err != nil {
			return nil, 0, err
		}
	
		totalLikesLastTwoWeeks, err := pr.GetTotalLikesLastTwoWeeks(allProduct[i].ID, "product")
		if err != nil {
			return nil, 0, err
		}
	
		productVm.TotalLike = totalLikes
		productVm.TotalLike = totalLikesLastTwoWeeks

		if totalLikes == 0 {
			productVm.Like = nil
		}

		var filteredLikes []response.Like
   		 for _, like := range productVm.Like {
        if like.SourceStr == "product" {
            filteredLikes = append(filteredLikes, like)
        	}
    	}
   		productVm.Like = filteredLikes

		resAllProduct = append(resAllProduct, *productVm)
	}

	var totalItems int64
	query.Count(&totalItems)

	// sort by like
	sort.Slice(resAllProduct, func(i, j int) bool {
		return resAllProduct[i].TotalLike > resAllProduct[j].TotalLike
	})

	return resAllProduct, int(totalItems), nil
}

func (pr *productRepository) GetTotalLikes(SourceId uint, SourceStr string) (uint, error) {
    var count int64
    if err := pr.db.Model(&models.Like{}).Where("source_id = ? AND source_str = ?", SourceId, SourceStr).Count(&count).Error; err != nil {
        return 0, err
    }
    return uint(count), nil
}

func (pr *productRepository) GetTotalLikesLastTwoWeeks(SourceId uint, SourceStr string) (uint, error) {
    var count int64

    if err := pr.db.Model(&models.Like{}).
        Where("source_id = ? AND source_str = ?", SourceId, SourceStr).
        Count(&count).Error; err != nil {
        return 0, err
    }
    return uint(count), nil
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