package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type WishlistHandler struct {
	wishlistService services.IWishlistService
}

func NewWishlistHandler(iWishlistService services.IWishlistService) *WishlistHandler {
	return &WishlistHandler{wishlistService: iWishlistService}
}

func (wi *WishlistHandler) CreateWishlist(c echo.Context) error {
	var input request.Wishlist
	c.Bind(&input)
	res, err := wi.wishlistService.CreateWishlist(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (wi *WishlistHandler) GetAllWishlist(c echo.Context) error {
	nameFilter := c.QueryParam("name")

    res, err := wi.wishlistService.GetAllWishlist(nameFilter)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}

func (wi *WishlistHandler) GetWishlist(c echo.Context) error {
	id := c.Param("id")
	res, err := wi.wishlistService.GetWishlist(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (wi *WishlistHandler) UpdateWishlist(c echo.Context) error {
	id := c.Param("id")
	var input request.Wishlist
	c.Bind(&input)

	res, err := wi.wishlistService.UpdateWishlist(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (wi *WishlistHandler) DeleteWishlist(c echo.Context) error {
	id := c.Param("id")
	res, err := wi.wishlistService.DeleteWishlist(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}