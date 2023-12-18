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

func RatingRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewRatingRepository(db)
	service := services.NewRatingService(repository)
	handler := handler.NewRatingHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/rating", handler.CreateRating)
	eJwt.GET("/rating", handler.GetAllRating)
	eJwt.GET("/rating/:id", handler.GetRating)
	eJwt.PUT("/rating/:id", handler.UpdateRating)
	eJwt.DELETE("/rating/:id", handler.DeleteRating)
}