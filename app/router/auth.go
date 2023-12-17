package routes

import (
	"lokasani/features/handler"
	"lokasani/features/services"
	"lokasani/features/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthGoogleRoute(e *echo.Echo, db *gorm.DB) {
	RoleRepo := repositories.NewRoleRepository(db)
	UserRepo := repositories.NewUsersRepository(db)
	AuthService := services.NewAuthService(RoleRepo, UserRepo)
	AuthHandler := handler.NewAuthHandler(*AuthService)

	e.GET("/google-auth", AuthHandler.OauthGoogleHandler)
	e.GET("/callback-google-auth", AuthHandler.OauthCallbackGoogleHandler)
}