package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func CommentRoute(e *echo.Echo, db *gorm.DB) {
	repositories := repositories.NewCommentRepository(db)
	service := services.NewCommentService(repositories)
	handler := handler.NewCommentHandler(service)

	e.POST("/comment", handler.CreateComment)
	e.GET("/comment", handler.GetAllComment)
	e.GET("/comment/:id", handler.GetComment)
	e.PUT("/comment/:id", handler.UpdateComment)
	e.DELETE("/comment/:id", handler.DeleteComment)
}