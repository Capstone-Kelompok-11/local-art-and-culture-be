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

func ArticleRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewArticleRepository(db)
	service := services.NewArticleService(repository)
	handler := handler.NewArticleHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	//admin can access
	eJwt.POST("/article", handler.CreateArticle)
	eJwt.PUT("/article/:id", handler.UpdateArticle)
	eJwt.DELETE("/article/:id", handler.DeleteArticle)

	//
	eJwt.GET("/article", handler.GetAllArticle)
	eJwt.GET("/article/:id", handler.GetArticle)
	eJwt.GET("/article/trending", handler.GetTrendingArticle)
}
