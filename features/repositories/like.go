package repositories

import (
	"fmt"
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	GetAllLike(articleId string) ([]response.Like, error)
	UpdateLike(input request.Like) (response.Like, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{db}
}

// func (co *likeRepository) CreateLike(data *request.Like) (response.Like, error) {
// 	dataLike := domain.ConvertFromLikeReqToModel(*data)
// 	err := co.db.Create(&data.Like).Error
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	err = co.db.First(&dataLike, "id = ?", dataLike.ID).Error
// 	return *domain.ConvertFromModelToLikeRes(*dataLike), nil
// }

func (co *likeRepository) GetAllLike(articleId string) ([]response.Like, error) {
	var allLike []models.Like
	var resAllLike []response.Like
	err := co.db.Where("article_id = ?", articleId).Where("active = ?", true).Find(&allLike).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	fmt.Println(allLike)

	for i := 0; i < len(allLike); i++ {
		likeVm := domain.ConvertFromModelToLikeRes(allLike[i])
		resAllLike = append(resAllLike, *likeVm)
	}
	return resAllLike, nil
}

// func (co *likeRepository) GetLike(id string) (response.Like, error) {
// 	var likeData models.Like
// 	err := co.db.First(&likeData, "id = ?", id).Error
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	return *domain.ConvertFromModelToLikeRes(likeData), nil
// }

func (co *likeRepository) UpdateLike(input request.Like) (response.Like, error) {
	likeData := *domain.ConvertFromLikeReqToModel(input)
	err := co.db.First(&likeData, "article_id = ? AND user_id = ? ", input.ArticleId, input.UserId).Error

	if input.Active != likeData.Active {
		likeData.Active = input.Active
	}

	if err = co.db.Save(&likeData).Error; err != nil {
		return response.Like{}, err
	}
	return *domain.ConvertFromModelToLikeRes(likeData), nil
}

// func (co *likeRepository) DeleteLike(id string) (response.Like, error) {
// 	likeData := models.Like{}
// 	res := response.Like{}
// 	find := co.db.First(&likeData, "id = ?", id).Error
// 	if find == nil {
// 		res = *domain.ConvertFromModelToLikeRes(likeData)
// 	}

// 	err := co.db.Delete(&likeData, "id = ?", id).Error
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	return res, nil
// }