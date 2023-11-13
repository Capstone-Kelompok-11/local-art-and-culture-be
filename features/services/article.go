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
		return errors.ERR_NAME_IS_EMPTY, response.Article{}
	}

	err, res := as.articleRepository.CreateArticle(data)
	if err != nil {
		return errors.ERR_REGISTER_USER_DATABASE, response.Article{}
	}

	return nil, res
}
