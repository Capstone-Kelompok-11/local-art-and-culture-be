package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	CreateArticle(data *request.Article) (error, response.Article)
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
