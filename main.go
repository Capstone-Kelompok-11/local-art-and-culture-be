package main

import (
	"fmt"
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/mysql"
	routes "lokasani/app/router"
)

func main() {
	appConfig, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)
	e := routes.Route(db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.APP_PORT)))
}
