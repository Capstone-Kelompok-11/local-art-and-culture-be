package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type ArticleHandler struct {
	articleService services.IArticleService
}

func NewArticleHandler(IArticleService services.IArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: IArticleService}
}

func (ah *ArticleHandler) CreateArticle(c echo.Context) error {
	var input request.Article
	c.Bind(&input)

	res, err := ah.articleService.CreateArticle(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ArticleHandler) GetAllArticle(c echo.Context) error {
	res, err := ah.articleService.GetAllArticle()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ArticleHandler) GetArticle(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.articleService.GetArticle(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ArticleHandler) UpdateArticle(c echo.Context) error {
	id := c.Param("id")
	var input request.Article
	c.Bind(&input)
	res, err := ah.articleService.UpdateArticle(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ArticleHandler) DeleteArticle(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.articleService.DeleteArticle(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}