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

type IArticleRepository interface {
	CreateArticle(data *request.Article) (error, response.Article)
	GetAllArticle() (error, []response.Article)
	GetArticle(id string) (error, response.Article)
	UpdateArticle(id string, input request.Article) (error, response.Article)
	DeleteArticle(id string) (error, response.Article)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) CreateArticle(data *request.Article) (error, response.Article) {
	dataArticle := domain.ConvertFromArticleReqToModel(*data)
	err := ar.db.Create(&dataArticle).Error
	if err != nil {
		return err, response.Article{}
	}
	return nil, *domain.ConvertFromModelToArticleRes(*dataArticle)
}

func (ar *articleRepository) GetAllArticle() (error, []response.Article) {
	var allArticle []models.Article
	var resAllArticle []response.Article
	err := ar.db.Find(&allArticle).Error
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}

	fmt.Println(allArticle)
	for i := 0; i < len(allArticle); i++ {
		articleVm := domain.ConvertFromModelToArticleRes(allArticle[i])
		resAllArticle = append(resAllArticle, *articleVm)
	}
	return nil, resAllArticle
}

func (ar *articleRepository) GetArticle(id string) (error, response.Article) {
	var articleData models.Article
	err := ar.db.First(&articleData, "id = ?", id).Error
	if err != nil {
		return err, response.Article{}
	}
	return nil, *domain.ConvertFromModelToArticleRes(articleData)
}

func (ar *articleRepository) UpdateArticle(id string, input request.Article) (error, response.Article) {
	articleData := models.Article{}
	err := ar.db.First(&articleData, "id = ?", id).Error
	if err != nil {
		return err, response.Article{}
	}

	if input.Title != "" {
		articleData.Title = input.Title
	} else if input.Content != "" {
		articleData.Content = input.Content
	}

	if err = ar.db.Save(&articleData).Error; err != nil {
		return err, response.Article{}
	}
	return nil, *domain.ConvertFromModelToArticleRes(articleData)
}

func (ar *articleRepository) DeleteArticle(id string) (error, response.Article) {
	articleData := models.Article{}
	res := response.Article{}
	find := ar.db.First(&articleData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToArticleRes(articleData)
	}
	err := ar.db.Delete(&articleData, "id = ?", id).Error
	if err != nil {
		return err, response.Article{}
	}
	return nil, res
}