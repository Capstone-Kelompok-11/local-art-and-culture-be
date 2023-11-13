package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewUsersRepository(db)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	e.POST("/users/register", handler.RegisterUsers)
	e.POST("/users/login", handler.LoginUsers)
	e.GET("/users", handler.GetAllUser)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
}