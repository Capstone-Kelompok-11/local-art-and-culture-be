package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	consts "lokasani/helpers/const"
	"lokasani/helpers/middleware"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleService services.IArticleService
}

func NewArticleHandler(IArticleService services.IArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: IArticleService}
}

func (ah *ArticleHandler) CreateArticle(c echo.Context) error {
	_, role, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if role != consts.AdminRole {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}
	
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
	page, _ := strconv.Atoi(c.QueryParam("page"))
    pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

    if page <= 0 {
        page = 1
    }

    if pageSize <= 0 {
        pageSize = 10
    }

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
	_, roleId, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if roleId != consts.AdminRole {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}

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
	_, roleId, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if roleId != consts.AdminRole {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}
	
    delete := c.Param("id")

    res, err := ah.articleService.DeleteArticle(delete)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}
