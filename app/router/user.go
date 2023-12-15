package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"lokasani/helpers/middleware"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewUsersRepository(db)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWTMiddleware())

	//access without token
	e.POST("/users/register", handler.RegisterUsers)
	e.POST("/users/login", handler.LoginUsers)

	// superadmin can access
	eJwt.GET("/users", handler.GetAllUser)
	eJwt.GET("/users/:id", handler.GetUser)
	eJwt.PUT("/users/:id", handler.UpdateUser)
	eJwt.DELETE("/users/:id", handler.DeleteUser)
}