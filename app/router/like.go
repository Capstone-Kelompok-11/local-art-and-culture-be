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

func LikeRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewLikeRepository(db)
	service := services.NewLikeService(repository)
	handler := handler.NewLikeHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	
	eJwt.PUT("/like", handler.UpdateLike)
	eJwt.GET("/like/:sourceId", handler.GetAllLike)
}
