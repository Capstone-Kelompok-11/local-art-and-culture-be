package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
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


func (ah *ArticleHandler) GetTrendingArticle(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	page, pageSize := 1, 10

	res, allItems, err := ah.articleService.GetTrendingArticle(nameFilter, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := ah.articleService.CalculatePaginationValues(page, pageSize, allItems)
	nextPage := ah.articleService.GetNextPage(currentPage, allPages)
	prevPage := ah.articleService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"data": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":  allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
}

func (ah *ArticleHandler) GetAllArticle(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	page, pageSize := 1, 10

	res, allItems, err := ah.articleService.GetAllArticle(nameFilter, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := ah.articleService.CalculatePaginationValues(page, pageSize, allItems)
	nextPage := ah.articleService.GetNextPage(currentPage, allPages)
	prevPage := ah.articleService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"data": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":  allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
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
    delete := c.Param("id")

    res, err := ah.articleService.DeleteArticle(delete)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}
