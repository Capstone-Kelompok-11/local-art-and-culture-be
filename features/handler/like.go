package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type LikeHandler struct {
	likeService services.ILikeService
}

func NewLikeHandler(ilikeService services.ILikeService) *LikeHandler {
	return &LikeHandler{likeService: ilikeService}
}


func (co *LikeHandler) GetAllLike(c echo.Context) error {
	id := c.Param("sourceId")
	res, err := co.likeService.GetAllLike(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *LikeHandler) UpdateLike(c echo.Context) error {
	var input request.Like
	c.Bind(&input)

	res, err := co.likeService.UpdateLike(input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
