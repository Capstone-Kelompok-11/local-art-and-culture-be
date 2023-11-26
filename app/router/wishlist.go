package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func WishlistRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewWishlistRepository(db)
	service := services.NewWishlistService(repository)
	handler := handler.NewWishlistHandler(service)

	e.POST("/wishlist", handler.CreateWishlist)
	e.GET("/wishlist", handler.GetAllWishlist)
	e.GET("/wishlist/:id", handler.GetWishlist)
	e.PUT("/wishlist/:id", handler.UpdateWishlist)
	e.DELETE("/wishlist/:id", handler.DeleteWishlist)
}