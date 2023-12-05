package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

type RatingHandler struct {
	ratingService services.IRatingService
}

func NewRatingHandler(iRatingService services.IRatingService) *RatingHandler {
	return &RatingHandler{ratingService: iRatingService}
}

func (ra *RatingHandler) CreateRating(c echo.Context) error {
	var input request.Rating
	c.Bind(&input)
	res, err := ra.ratingService.CreateRating(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ra *RatingHandler) GetAllRating(c echo.Context) error {
	nameFilter := c.QueryParam("name")

    res, err := ra.ratingService.GetAllRating(nameFilter)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}

func (ra *RatingHandler) GetRating(c echo.Context) error {
	id := c.Param("id")
	res, err := ra.ratingService.GetRating(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ra *RatingHandler) UpdateRating(c echo.Context) error {
	id := c.Param("id")
	var input request.Rating
	c.Bind(&input)

	res, err := ra.ratingService.UpdateRating(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ra *RatingHandler) DeleteRating(c echo.Context) error {
	id := c.Param("id")
	res, err := ra.ratingService.DeleteRating(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}