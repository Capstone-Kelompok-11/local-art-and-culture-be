package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func LikeRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewLikeRepository(db)
	service := services.NewLikeService(repository)
	handler := handler.NewLikeHandler(service)

	e.PUT("/like", handler.UpdateLike)
	e.GET("/like/:articleId", handler.GetAllLike)
}
