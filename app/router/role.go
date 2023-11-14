package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RoleRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repository)
	handler := handler.NewRoleHandler(service)

	e.POST("/role", handler.CreateRole)
	e.GET("/role", handler.GetAllRole)
	e.GET("/role/:id", handler.GetRole)
	e.PUT("/role/:id", handler.UpdateRole)
	e.DELETE("/role/:id", handler.DeleteRole)
}
