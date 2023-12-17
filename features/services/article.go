package services

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
	"math"
)

type IArticleService interface {
	CreateArticle(data *request.Article) (response.Article, error)
	GetTrendingArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error)
	GetAllArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error)
	GetArticle(id string) (response.Article, error)
	UpdateArticle(id string, input request.Article) (response.Article, error)
	DeleteArticle(id string) (response.Article, error)
	CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	GetNextPage(currentPage, allPages int) int
	GetPrevPage(currentPage int) int
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
	if data.Status == "" {
		return response.Article{}, errors.ERR_STATUS_IS_EMPTY
	}

	res, err := as.articleRepository.CreateArticle(data)
	if err != nil {
		return response.Article{}, errors.ERR_CREATE_ARTICLE_DATABASE
	}
	fmt.Println(data)
	return res, nil
}

func (as *ArticleService) GetTrendingArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error) {
	res, allItems, err := as.articleRepository.GetTrendingArticle(nameFilter, page, pageSize)
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}
	return res, allItems, nil
}

func (as *ArticleService) GetAllArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error) {
	res, allItems, err := as.articleRepository.GetAllArticle(nameFilter, page, pageSize)
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}
	return res, allItems, nil
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

func (pr *ArticleService) CalculatePaginationValues(page, pageSize, allItmes int) (int, int) {
	pageInt := page
	if pageInt <= 0 {
		pageInt = 1
	}

	allPages := int(math.Ceil(float64(allItmes) / float64(pageSize)))

	if pageInt > allPages {
		pageInt = allPages
	}

	return pageInt, allPages
}

func (pr *ArticleService) GetNextPage(currentPage, allPages int) int {
	if currentPage < allPages {
		return currentPage + 1
	}
	return allPages
}

func (pr *ArticleService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}
