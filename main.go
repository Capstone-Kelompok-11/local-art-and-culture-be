package main

import (
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	_, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)
	
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  		AllowOrigins: []string{"*", "*"},
  		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	routes.Route(e, db)
	e.Logger.Fatal(e.Start(":80"))
}