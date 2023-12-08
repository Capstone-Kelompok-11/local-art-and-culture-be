package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentService services.ICommentService
}

func NewCommentHandler(iCommentService services.ICommentService) *CommentHandler{
	return &CommentHandler{commentService: iCommentService}
}

func (co *CommentHandler) CreateComment(c echo.Context) error {
	var input request.Comment
	c.Bind(&input)
	res, err := co.commentService.CreateComment(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *CommentHandler) GetAllComment(c echo.Context) error {
	res, err := co.commentService.GetAllComment()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *CommentHandler) GetComment(c echo.Context) error {
	id := c.Param("id")
	res, err := co.commentService.GetComment(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *CommentHandler) UpdateComment(c echo.Context) error {
	id := c.Param("id")
	var input request.Comment
	c.Bind(&input)

	res, err := co.commentService.UpdateComment(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (co *CommentHandler) DeleteComment(c echo.Context) error {
	id := c.Param("id")
	res, err := co.commentService.DeleteComment(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}