package routes

import ( 
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func CreatorRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewCreatorRepository(db)
	service := services.NewCreatorService(repository)
	handler := handler.NewCreatorHandler(service)

	e.POST("/creator", handler.CreateCreator)
	e.GET("/creator", handler.GetAllCreator)
	e.GET("/creator/:id", handler.GetCreator)
	e.PUT("/creator/:id", handler.UpdateCreator)
	e.DELETE("/creator/:id", handler.DeleteCreator)
}
	