package routes

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) *echo.Echo {
	godotenv.Load(".env")
	e := echo.New()
	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	RoleRoute(e, db)
	return e
}
