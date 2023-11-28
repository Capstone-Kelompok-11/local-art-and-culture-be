package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func FilesRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewFilesRepository(db)
	service := services.NewFilesService(repository)
	handler := handler.NewFilesHandler(service)

	e.POST("/files", handler.CreateFiles)
	e.GET("/files", handler.GetAllFiles)
	e.GET("/files/:id", handler.GetFiles)
	e.PUT("/files/:id", handler.UpdateFiles)
	e.DELETE("/files/:id", handler.DeleteFiles)
}
