package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ArticleRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewArticleRepository(db)
	service := services.NewArticleService(repository)
	handler := handler.NewArticleHandler(service)

	e.POST("/article", handler.CreateArticle)
}
