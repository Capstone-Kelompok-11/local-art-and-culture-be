package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"os"

	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewUsersRepository(db)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	//access without token
	e.POST("/users/register", handler.RegisterUsers)
	e.POST("/users/login", handler.LoginUsers)

	//access with token
	e.GET("/users", handler.GetAllUser)
	e.GET("/users/:id", handler.GetUser)
	eJwt.PUT("/users/:id", handler.UpdateUser)
	eJwt.DELETE("/users/:id", handler.DeleteUser)
}