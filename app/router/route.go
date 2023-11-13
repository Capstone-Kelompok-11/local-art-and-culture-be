package routes

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) *echo.Echo {
	godotenv.Load()
	e := echo.New()
	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))
	AdminRoute(e, db)
	ArticleRoute(e, db)
	return e
}