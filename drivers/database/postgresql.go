package database

import (
	"fmt"

	"github.com/satryanararya/go-chefbot/entities"
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

	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{}, 
		&entities.UserFoodPreference{}, 
		&entities.UserCookingSkill{},
		&entities.Ingredient{},
		&entities.UserAllergies{},
		&entities.Recipe{},
		&entities.ExtendedIngredient{},
		&entities.RatingReview{},
		&entities.FavoriteRecipe{},
		&entities.Recommendation{},
	)
}