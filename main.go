package main

import (
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)

	e := routes.Route(db)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  		AllowOrigins: []string{"*", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
  		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.Start(":80"))
}
