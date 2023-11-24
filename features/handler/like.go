package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/labstack/echo"
)

type LikeHandler struct {
	likeService services.likeService
}

func NewLikeHandler(ilikeService services.likeService) *LikeHandler {
	return &LikeHandler{likeService: ilikeService}
}

func (co *LikeHandler) CreateLike(c echo.Context) error {
	var input request.Like
	c.Bind(&input)
	res, err := co.likeService.CreateComment(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *LikeHandler) GetAllLike(c echo.Context) error {
	res, err := co.likeService.GetAllLike()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *LikeHandler) GetLike(c echo.Context) error {
	id := c.Param("id")
	res, err := co.likeService.GetComment(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *LikeHandler) UpdateLike(c echo.Context) error {
	id := c.Param("id")
	var input request.Like
	c.Bind(&input)

	res, err := co.likeService.UpdateLike(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *LikeHandler) DeleteComment(c echo.Context) error {
	id := c.Param("id")
	res, err := co.likeService.DeleteComment(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
