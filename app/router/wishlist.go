package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func WishlistRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewWishlistRepository(db)
	service := services.NewWishlistService(repository)
	handler := handler.NewWishlistHandler(service)

	
	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/wishlist", handler.CreateWishlist)
	eJwt.GET("/wishlist", handler.GetAllWishlist)
	eJwt.GET("/wishlist/:id", handler.GetWishlist)
	eJwt.PUT("/wishlist/:id", handler.UpdateWishlist)
	eJwt.DELETE("/wishlist/:id", handler.DeleteWishlist)
}