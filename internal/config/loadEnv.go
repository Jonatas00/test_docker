package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOSTNAME string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string

	APP_PORT string
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	DB_HOSTNAME = os.Getenv("DB_HOSTNAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")

	APP_PORT = os.Getenv("APP_PORT")

	return nil
}
