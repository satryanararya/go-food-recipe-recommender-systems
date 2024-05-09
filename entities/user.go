package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                 int64  `gorm:"primaryKey"`
	Username           string `gorm:"unique;not null"`
	Email              string `gorm:"unique;not null"`
	Password           string
	Token              string
	UserFoodPreference UserFoodPreference 
	UserCookingSkill   UserCookingSkill   
	UserAllergies      []UserAllergies    
	CreatedAt          time.Time          `gorm:"autoCreateTime"`
	UpdatedAt          time.Time          `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt     `gorm:"index"`
}

// Additional user information - User Profile
type UserFoodPreference struct {
	gorm.Model
	UserID             int64 `gorm:"not null"`
	DietaryRestriction *string
	ReligiousReason    *string
}

type UserCookingSkill struct {
	gorm.Model
	UserID              int64 `gorm:"not null"`
	ExperienceYears     string
	TimeCommitment      string
	RecipeComplexity    string
	IngredientDiversity string
}

// TODO: Ngulik connect ke Ingredients
type UserAllergies struct {
	gorm.Model
	UserID       int64 `gorm:"not null"`
	IngredientID *int64
}
