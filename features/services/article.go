package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IArticleService interface {
	CreateArticle(data *request.Article) (error, response.Article)
	GetAllArticle() (error, []response.Article)
	GetArticle(id string) (error, response.Article)
	UpdateArticle(id string, input request.Article) (error, response.Article)
	DeleteArticle(id string) (error, response.Article)
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

func (as *ArticleService) GetAllArticle() (error, []response.Article) {
	err, res := as.articleRepository.GetAllArticle()
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}
	return nil, res
}

func (as *ArticleService) GetArticle(id string) (error, response.Article) {
	if id == "" {
		return errors.ERR_GET_BAD_REQUEST_ID, response.Article{}
	}

	err, res := as.articleRepository.GetArticle(id)
	if err != nil {
		return err, response.Article{}
	}
	return nil, res
}

func (as *ArticleService) UpdateArticle(id string, input request.Article) (error, response.Article) {
	if id == "" {
		return errors.ERR_GET_BAD_REQUEST_ID, response.Article{}
	}

	err, res := as.articleRepository.UpdateArticle(id, input)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.Article{}
	}
	return nil, res
}

func (as *ArticleService) DeleteArticle(id string) (error, response.Article) {
	if id == "" {
		return errors.ERR_GET_BAD_REQUEST_ID, response.Article{}
	}

	err, res := as.articleRepository.DeleteArticle(id)
	if err != nil {
		return errors.ERR_DELETE, response.Article{}
	}
	return nil, res
}