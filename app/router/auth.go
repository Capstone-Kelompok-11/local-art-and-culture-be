package routes

import (
	"lokasani/features/handler"
	"lokasani/features/services"
	"lokasani/features/repositories"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthGoogleRoute(e *echo.Group, db *gorm.DB, validate *validator.Validate) {
	RoleRepo := repositories.NewRoleRepository(db)
	UserRepo := repositories.NewUsersRepository(db)
	AuthService := services.NewAuthService(RoleRepo, UserRepo, validate)
	AuthHandler := handler.NewAuthHandler(*AuthService)

	v1 := e.Group("")
	v1.GET("/google-auth", AuthHandler.OauthGoogleHandler)
	v1.GET("/callback-google-auth", AuthHandler.OauthCallbackGoogleHandler)
}