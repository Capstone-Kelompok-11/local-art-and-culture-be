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

func AdminRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewAdminRepository(db)
	service := services.NewAdminService(repository)
	handler := handler.NewAdminHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	//access without token
	e.POST("/admin/register", handler.RegisterAdmin)
	e.POST("/admin/login", handler.LoginAdmin)

	//access with token
	eJwt.GET("/admin", handler.GetAllAdmin)
	eJwt.GET("/admin/:id", handler.GetAdmin)
	eJwt.PUT("/admin/:id", handler.UpdateAdmin)
	eJwt.DELETE("/admin/:id", handler.DeleteAdmin)
}