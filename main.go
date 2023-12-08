package main

import (
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)

	e := routes.Route(db)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Logger.Fatal(e.Start(":80"))
}
