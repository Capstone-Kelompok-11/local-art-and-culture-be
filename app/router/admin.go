package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AdminRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewAdminRepository(db)
	service := services.NewAdminService(repository)
	handler := handler.NewAdminHandler(service)

	e.POST("/admin/register", handler.RegisterAdmin)
	e.POST("/admin/login", handler.LoginAdmin)
	e.GET("/admin", handler.GetAllAdmin)
}
