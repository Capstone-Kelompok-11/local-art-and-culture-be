package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	APP_PORT   int
	SECRET_KEY string
}

func LoadAPP() *AppConfig {
	var result = new(AppConfig)

	// godotenv.Load(".env")

	if v, found := os.LookupEnv("APP_PORT"); found {
		port, err := strconv.Atoi(v)
		if err != nil {
			logrus.Error("Config: invalid port value,", err.Error())
			return nil
		}
		result.APP_PORT = port
	}
	return result
}
