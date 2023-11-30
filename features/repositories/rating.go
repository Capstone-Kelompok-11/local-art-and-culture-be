package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IRatingRepository interface {
	CreateRating(data *request.Rating) (response.Rating, error)
	GetAllRating(nameFilter string) ([]response.Rating, error)
	GetRating(id string) (response.Rating, error)
	UpdateRating(id string, input request.Rating) (response.Rating, error)
	DeleteRating(id string) (response.Rating, error)
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *ratingRepository {
	return &ratingRepository{db}
}

func (ra *ratingRepository) CreateRating(data *request.Rating) (response.Rating, error) {
	dataRating := domain.ConvertFromRatingReqToModel(*data)
	err := ra.db.Create(&dataRating).Error
	if err != nil {
		return response.Rating{}, err
	}
	err = ra.db.Preload("Users").Preload("Product").Preload("Product.Category").First(&dataRating, "id = ?", dataRating.ID).Error
	return *domain.ConvertFromModelToRatingRes(*dataRating), nil
}

func (ra *ratingRepository) GetAllRating(nameFilter string) ([]response.Rating, error) {
	var allRating []models.Rating
	var resAllRating []response.Rating

	query := ra.db.Preload("Users").Preload("Product").Preload("Product.Category")

	if nameFilter != "" {
		query = query.Where("rating LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allRating).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allRating); i++ {
		RatingVm := domain.ConvertFromModelToRatingRes(allRating[i])
		resAllRating = append(resAllRating, *RatingVm)
	}

	return resAllRating, nil
}

func (ra *ratingRepository) GetRating(id string) (response.Rating, error) {
	var ratingData models.Rating
	err := ra.db.Preload("Users").Preload("Product").Preload("Product.Category").First(&ratingData, "id = ?", id).Error

	if err != nil {
		return response.Rating{}, err
	}
	return *domain.ConvertFromModelToRatingRes(ratingData), nil
}

func (ra *ratingRepository) UpdateRating(id string, input request.Rating) (response.Rating, error) {
	ratingData := models.Rating{}
	err := ra.db.Preload("Users").Preload("Product").Preload("Product.Category").First(&ratingData, "id = ?", id).Error
	if err != nil {
		return response.Rating{}, err
	}

	if input.Rating != 0 {
		ratingData.Rating = input.Rating
	}
	if input.Ulasan != "" {
		ratingData.Ulasan = input.Ulasan
	}

	if err = ra.db.Save(&ratingData).Error; err != nil {
		return response.Rating{}, err
	}
	return *domain.ConvertFromModelToRatingRes(ratingData), nil
}

func (ra *ratingRepository) DeleteRating(id string) (response.Rating, error) {
	ratingData := models.Rating{}
	res := response.Rating{}
	find := ra.db.Preload("Users").Preload("Product").Preload("Product.Category").First(&ratingData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToRatingRes(ratingData)
	}

	err := ra.db.Delete(&ratingData, "id = ?", id).Error
	if err != nil {
		return response.Rating{}, err
	}
	return res, nil
}