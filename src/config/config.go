package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ConfigOptions struct {
	ENV                 string
	SERVER              string
	PORT                string
	API_VERSION         string
	EXPIRED_CONNECTIONS uint
}

var Config ConfigOptions = ConfigOptions{}

func LoadConfig() error {
	// Memuat variabel lingkungan dari berkas .env
	if err := godotenv.Load(); err != nil {
		return err
	}

	Config.ENV = os.Getenv("GO_ENV")

	if strings.ToLower(Config.ENV) == "development" {
		conf, errLoadConf := LoadEnvConfig(Config.ENV)
		if errLoadConf != nil {
			return errLoadConf
		}
		Config = conf
	}
	if strings.ToLower(Config.ENV) == "staging" {
		conf, errLoadConf := LoadEnvConfig(Config.ENV)
		if errLoadConf != nil {
			return errLoadConf
		}
		Config = conf
	}
	if strings.ToLower(Config.ENV) == "production" {
		conf, errLoadConf := LoadEnvConfig(Config.ENV)
		if errLoadConf != nil {
			return errLoadConf
		}
		Config = conf
	}

	return nil
}