package entities

import "github.com/google/uuid"

// import "mime/multipart"

type Recipe struct {
	ID                int64 `gorm:"primaryKey" json:"id"`
	UserID            uuid.UUID `json:"user_id"`
	Title             string `gorm:"not null" json:"title" validate:"required,min=2,max=50" form:"title"`
	Image             string `json:"image"`
	SourceName        string `json:"source_name"`
	CookingMinutes    int64 `json:"cooking_minutes"`
	RecipeIngredients []RecipeIngredient `gorm:"foreignKey:RecipeID" json:"recipe_ingredients"`
	PricePerServing   float64 `json:"price_per_serving"`
	ReadyInMinutes    int `json:"ready_in_minutes"`
	Servings          int `json:"servings"`
	Diets             string `json:"diets"`
	IsSustainable     bool `json:"is_sustainable"`
	Instruction       string `json:"instruction" validate:"required,min=2,max=500" form:"instruction"`
	RatingReview      *[]RatingReview   `gorm:"foreignKey:RecipeID" json:"rating_review"`
	FavoriteRecipe    *[]FavoriteRecipe `gorm:"foreignKey:RecipeID" json:"favorite_recipe"`
}

type RecipeIngredient struct {
	ID       int64 `gorm:"primaryKey"`
	RecipeID int64
	Name     string `gorm:"not null"`
	Quantity float64
	Unit     string
}
