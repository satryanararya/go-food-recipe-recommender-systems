package entities

// import "mime/multipart"

type Recipe struct {
	ID                  int64 `gorm:"primaryKey"`
	UserID              int64
	Title               string `gorm:"not null"`
	Image               string
	SourceName          string
	CookingMinutes      int64
	ExtendedIngredients []ExtendedIngredient `gorm:"foreignKey:RecipeID"`
	PricePerServing     float64
	ReadyInMinutes      int
	Servings            int
	HealthScore         float64
	Diets               string
	IsSustainable       bool
	Instruction         string
	RatingReview		*[]RatingReview `gorm:"foreignKey:RecipeID"`
}

type ExtendedIngredient struct {
	ID       int64 `gorm:"primaryKey"`
	RecipeID int64
	Name     string `gorm:"not null"`
	Amount   float64
	Unit     string
}
