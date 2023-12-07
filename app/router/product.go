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

func ProductRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewProductRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	service := services.NewProductService(repository, categoryRepository)
	handler := handler.NewProductHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/product", handler.CreateProduct)
	eJwt.GET("/product", handler.GetAllProduct)
	eJwt.GET("/product/trending", handler.GetTrendingProduct)
	eJwt.GET("/product/:id", handler.GetProduct)
	eJwt.PUT("/product/:id", handler.UpdateProduct)
	eJwt.DELETE("/product/:id", handler.DeleteProduct)
}
