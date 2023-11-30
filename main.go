package main

import (
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"

	"github.com/labstack/echo/middleware"
)

func main() {
	_, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)
	e := routes.Route(db)
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	  }))
	  
	e.Logger.Fatal(e.Start(":8080"))
}
