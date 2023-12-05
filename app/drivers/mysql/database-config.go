package mysql

import (
	"fmt"
	"lokasani/app/drivers/migrations"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_PORT     int
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func LoadDB() *DBConfig {
	var result = new(DBConfig)

	godotenv.Load(".env")
	if v, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(v)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		result.DB_PORT = port
	}

	if v, found := os.LookupEnv("DB_HOST"); found {
		result.DB_HOST = v
	}

	if v, found := os.LookupEnv("DB_USER"); found {
		result.DB_USER = v
	}

	if v, found := os.LookupEnv("DB_PASSWORD"); found {
		result.DB_PASSWORD = v
	}

	if v, found := os.LookupEnv("DB_NAME"); found {
		result.DB_NAME = v
	}
	return result
}

func StartDB(cfg *DBConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	migrations.InitMigrate(DB)
	return DB
}