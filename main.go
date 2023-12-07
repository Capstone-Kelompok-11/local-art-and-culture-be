package main

import (
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"
)

func main() {
	_, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)
	e := routes.Route(db)

	e.Logger.Fatal(e.Start(":80"))
}
