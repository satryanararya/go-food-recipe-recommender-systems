package entities

import "time"

type FavoriteRecipe struct {
	ID          int64 `gorm:"primaryKey"`
	UserID      int64
	RecipeID    int64
	RecipeTitle string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
