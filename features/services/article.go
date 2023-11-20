package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IArticleService interface {
	CreateArticle(data *request.Article) (response.Article, error)
	GetAllArticle() ([]response.Article, error)
	GetArticle(id string) (response.Article, error)
	UpdateArticle(id string, input request.Article) (response.Article, error)
	DeleteArticle(id string) (response.Article, error)
	//GetAdminWithArticles(adminID uint) (*response.Article, error)
}

type ArticleService struct {
	articleRepository repositories.IArticleRepository
}

func NewArticleService(repo repositories.IArticleRepository) *ArticleService {
	return &ArticleService{articleRepository: repo}
}

func (as *ArticleService) CreateArticle(data *request.Article) (response.Article, error) {
	if data.Title == "" {
		return response.Article{}, errors.ERR_NAME_IS_EMPTY
	}

	res, err := as.articleRepository.CreateArticle(data)
	if err != nil {
		return response.Article{}, errors.ERR_REGISTER_USER_DATABASE
	}

	return res, nil
}

func (as *ArticleService) GetAllArticle() ([]response.Article, error) {
	res, err := as.articleRepository.GetAllArticle()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (as *ArticleService) GetArticle(id string) (response.Article, error) {
	if id == "" {
		return response.Article{}, errors.ERR_GET_BAD_REQUEST_ID
	}

	res, err := as.articleRepository.GetArticle(id)
	if err != nil {
		return response.Article{}, err
	}
	return res, nil
}

func (as *ArticleService) UpdateArticle(id string, input request.Article) (response.Article, error) {
	if id == "" {
		return response.Article{}, errors.ERR_GET_BAD_REQUEST_ID
	}

	res, err := as.articleRepository.UpdateArticle(id, input)
	if err != nil {
		return response.Article{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (as *ArticleService) DeleteArticle(id string) (response.Article, error) {
	if id == "" {
		return response.Article{}, errors.ERR_GET_BAD_REQUEST_ID
	}

	res, err := as.articleRepository.DeleteArticle(id)
	if err != nil {
		return response.Article{}, errors.ERR_DELETE
	}
	return res, nil
}
