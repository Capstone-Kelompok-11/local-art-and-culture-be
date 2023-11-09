package config

import (
	"lokasani/app/drivers/mysql"

	"github.com/sirupsen/logrus"
)

func InitConfig() (*AppConfig, *mysql.DBConfig) {
	db := mysql.LoadDB()
	app := LoadAPP()

	if db == nil || app == nil {
		logrus.Fatal("Config : cannot start program, failed to load configuration")
		return nil, nil
	}
	return app, db
}
