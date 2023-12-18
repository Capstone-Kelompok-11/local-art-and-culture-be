package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewCategoryRepository(db)
	service := services.NewCategoryService(repository)
	handler := handler.NewCategoryHandler(*service)

	e.POST("/category", handler.CreateCategory)
	e.GET("/category", handler.GetAllCategory)
	e.GET("/category/:id", handler.GetCategory)
	e.GET("/category/:Type", handler.GetTypeCategory)
	e.PUT("/category/:id", handler.UpdateCategory)
	e.DELETE("/category/:id", handler.DeleteCategory)
}