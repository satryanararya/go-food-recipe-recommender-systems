package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB_HOST,
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}