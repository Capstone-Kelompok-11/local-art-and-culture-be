package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RatingRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewRatingRepository(db)
	service := services.NewRatingService(repository)
	handler := handler.NewRatingHandler(service)

	e.POST("/rating", handler.CreateRating)
	e.GET("/rating", handler.GetAllRating)
	e.GET("/rating/:id", handler.GetRating)
	e.PUT("/rating/:id", handler.UpdateRating)
	e.DELETE("/rating/:id", handler.DeleteRating)
}