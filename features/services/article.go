package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IArticleService interface {
	CreateArticle(data *request.Article) (error, response.Article)
}

type ArticleService struct {
	articleRepository repositories.IArticleRepository
}

func NewArticleService(repo repositories.IArticleRepository) *ArticleService {
	return &ArticleService{articleRepository: repo}
}

func (as *ArticleService) CreateArticle(data *request.Article) (error, response.Article) {
	if data.Title == "" {
		return errors.ERR_TITLE_IS_EMPTY, response.Article{}
	}

	if data.Content == "" {
		return errors.ERR_CONTENT_IS_EMPTY, response.Article{}
	}

	err, res := as.articleRepository.CreateArticle(data)
	if err != nil {
		return errors.ERR_CREATE_ARTICLE_DATABASE, response.Article{}
	}

	return nil, res
}
