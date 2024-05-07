package config

import (
	"log"
	"os"

	"github.com/satryanararya/go-chefbot/drivers/database"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func InitConfigDB() database.Config {
	return database.Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
	}
}