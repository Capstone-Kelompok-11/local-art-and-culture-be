package main

import (
	"fmt"
	"lokasani/app/drivers/config"
	"lokasani/app/drivers/migrations"

	"github.com/labstack/echo/v4"
)

func main() {
	appConfig, dbConfig := config.InitConfig()
	migrations.InitMigrate(dbConfig)

	app := echo.New()
	
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appConfig.APP_PORT)))
}
