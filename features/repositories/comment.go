package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(data *request.Comment) (response.Comment, error)
	GetAllComment() ([]response.Comment, error)
	GetComment(id string) (response.Comment, error)
	UpdateComment(id string, input request.Comment) (response.Comment, error)
	DeleteComment(id string) (response.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository{
	return &commentRepository{db}
}

func (co *commentRepository) CreateComment(data *request.Comment) (response.Comment, error) {
	dataComment := domain.ConvertFromCommentReqToModel(*data)
	err := co.db.Create(&dataComment).Error
	if err != nil {
		return response.Comment{}, err
	}
	err = co.db.Preload("Article").First(&dataComment, "id = ?", dataComment.ID).Error
	return *domain.ConvertFromModelToCommentRes(*dataComment), nil
}

func (co *commentRepository) GetAllComment() ([]response.Comment, error) {
	var allComment []models.Comment
	var resAllComment []response.Comment
	err := co.db.Preload("Article").Find(&allComment).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allComment); i++ {
		commentVm := domain.ConvertFromModelToCommentRes(allComment[i])
		resAllComment = append(resAllComment, *commentVm)
	}
	return resAllComment, nil
}

func (co *commentRepository) GetComment(id string) (response.Comment, error) {
	var commentData models.Comment
	err := co.db.Preload("Article").First(&commentData, "id = ?", id).Error
	if err != nil {
		return response.Comment{}, err
	}
	return *domain.ConvertFromModelToCommentRes(commentData), nil
}

func (co *commentRepository) UpdateComment(id string, input request.Comment) (response.Comment, error) {
	commentData := models.Comment{}
	err := co.db.First(&commentData, "id = ?", id).Error
	if err != nil {
		return response.Comment{}, err
	}

	if input.Comment != "" {
		commentData.Comment = input.Comment
	}
	if input.ArticleId != 0 {
		commentData.ArticleId = input.ArticleId
	}

	if err = co.db.Save(&commentData).Error; err != nil {
		return response.Comment{}, err
	}
	return *domain.ConvertFromModelToCommentRes(commentData), nil
}

func (co *commentRepository) DeleteComment(id string) (response.Comment, error) {
	commentData := models.Comment{}
	res := response.Comment{}
	find := co.db.Preload("Article").First(&commentData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToCommentRes(commentData)
	}

	err := co.db.Delete(&commentData, "id = ?", id).Error
	if err != nil {
		return response.Comment{}, err
	}
	return res, nil
}