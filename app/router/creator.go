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

func CreatorRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewCreatorRepository(db)
	service := services.NewCreatorService(repository)
	handler := handler.NewCreatorHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/creator", handler.CreateCreator)
	eJwt.GET("/creator", handler.GetAllCreator)
	eJwt.GET("/creator/role/:role", handler.GetAllCreatorByRole)
	eJwt.GET("/creator/:id", handler.GetCreator)
	eJwt.PUT("/creator/:id", handler.UpdateCreator)
	eJwt.DELETE("/creator/:id", handler.DeleteCreator)
}