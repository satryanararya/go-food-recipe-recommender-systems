package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 uuid.UUID `gorm:"primaryKey"`
	Username           string    `gorm:"unique;not null"`
	Email              string    `gorm:"unique;not null"`
	Password           string
	UserFoodPreference UserFoodPreference `gorm:"foreignKey:UserID"`
	UserCookingSkill   UserCookingSkill   `gorm:"foreignKey:UserID"`
	UserAllergies      []UserAllergies    `gorm:"foreignKey:UserID"`
	Recipe             []Recipe           `gorm:"foreignKey:UserID"`
	RatingReview       *[]RatingReview    `gorm:"foreignKey:UserID"`
	FavoriteRecipe     *[]FavoriteRecipe  `gorm:"foreignKey:UserID"`
	Recommendations    *[]Recommendation
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

// Additional user information - User Profile
type UserFoodPreference struct {
	ID                 int64          `gorm:"primaryKey" json:"-"`
	UserID             uuid.UUID      `gorm:"not null" json:"-"`
	DietaryRestriction *string        `json:"dietary_restriction"`
	ReligiousReason    *string        `json:"religious_reason"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserCookingSkill struct {
	ID                  int64          `gorm:"primaryKey" json:"-"`
	UserID              uuid.UUID      `gorm:"not null" json:"-"`
	ExperienceYears     string         `json:"experience_years"`
	TimeCommitment      string         `json:"time_commitment"`
	RecipeComplexity    string         `json:"recipe_complexity"`
	IngredientDiversity string         `json:"ingredient_diversity"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
}

// TODO: Ngulik connect ke Ingredients
type UserAllergies struct {
	ID           int64     `gorm:"primaryKey" json:"-"`
	UserID       uuid.UUID `gorm:"not null" json:"-"`
	IngredientID int64     `json:"ingredient_id"`
	Ingredient   Ingredient
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
