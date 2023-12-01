package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ProductRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewProductRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	service := services.NewProductService(repository, categoryRepository)
	handler := handler.NewProductHandler(service)

	e.POST("/product", handler.CreateProduct)
	e.GET("/product", handler.GetAllProduct)
	e.GET("/product/:id", handler.GetProduct)
	e.PUT("/product/:id", handler.UpdateProduct)
	e.DELETE("/product/:id", handler.DeleteProduct)
}
