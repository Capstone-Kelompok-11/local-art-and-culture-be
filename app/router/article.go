package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"lokasani/helpers/middleware"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ArticleRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewArticleRepository(db)
	service := services.NewArticleService(repository)
	handler := handler.NewArticleHandler(service)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWTMiddleware())

	//admin can access
	eJwt.POST("/article", handler.CreateArticle)
	eJwt.PUT("/article/:id", handler.UpdateArticle)
	eJwt.DELETE("/article/:id", handler.DeleteArticle)

	//without token
	e.GET("/article", handler.GetAllArticle)
	e.GET("/article/:id", handler.GetArticle)
	e.GET("/article/trending", handler.GetTrendingArticle)
}
