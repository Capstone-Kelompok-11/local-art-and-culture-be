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
	// e.GET("/article", handler.GetAllArticle)
	e.GET("/article/trending", handler.GetTrendingArticle)
	e.GET("/article/:id", handler.GetArticle)
	e.PUT("/article/:id", handler.UpdateArticle)
	e.DELETE("/article/:id", handler.DeleteArticle)
}
