package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	GetAllLike(sourceId string) ([]response.Like, error)
	UpdateLike(input request.Like) (response.Like, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{db}
}

func (co *likeRepository) GetAllLike(sourceId string) ([]response.Like, error) {
	var allLike []models.Like
	var resAllLike []response.Like
	err := co.db.Where("source_id = ?", sourceId).Where("active = ?", true).Find(&allLike).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allLike); i++ {
		likeVm := domain.ConvertFromModelToLikeRes(allLike[i])
		resAllLike = append(resAllLike, *likeVm)
	}
	return resAllLike, nil
}

func (co *likeRepository) UpdateLike(input request.Like) (response.Like, error) {
	likeData := *domain.ConvertFromLikeReqToModel(input)
	err := co.db.First(&likeData, "source_id = ? AND user_id = ? AND source_str = ?", input.SourceId, input.UserId, input.SourceStr).Error

	if input.Active != likeData.Active {
		likeData.Active = input.Active
	}
	if input.SourceStr != "" {
		likeData.SourceStr = input.SourceStr
	}

	if err = co.db.Save(&likeData).Error; err != nil {
		return response.Like{}, err
	}
	return *domain.ConvertFromModelToLikeRes(likeData), nil
}
