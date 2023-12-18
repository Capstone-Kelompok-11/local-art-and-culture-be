package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"lokasani/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewProductRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	service := services.NewProductService(repository, categoryRepository)
	handler := handler.NewProductHandler(service)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWTMiddleware())

	//product creator can access
	eJwt.POST("/product", handler.CreateProduct)
	eJwt.PUT("/product/:id", handler.UpdateProduct)
	eJwt.DELETE("/product/:id", handler.DeleteProduct)

	//without token
	e.GET("/product", handler.GetAllProduct)
	e.GET("/product/trending", handler.GetTrendingProduct)
	e.GET("/product/:id", handler.GetProduct)
}
