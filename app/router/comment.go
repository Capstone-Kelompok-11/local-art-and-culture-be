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

func CommentRoute(e *echo.Echo, db *gorm.DB) {
	repositories := repositories.NewCommentRepository(db)
	service := services.NewCommentService(repositories)
	handler := handler.NewCommentHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	
	eJwt.POST("/comment", handler.CreateComment)
	eJwt.GET("/comment", handler.GetAllComment)
	eJwt.GET("/comment/:id", handler.GetComment)
	eJwt.PUT("/comment/:id", handler.UpdateComment)
	eJwt.DELETE("/comment/:id", handler.DeleteComment)
}