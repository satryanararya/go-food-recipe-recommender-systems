package entities

import (
	"time"

	"github.com/google/uuid"
)

type FavoriteRecipe struct {
	ID          int64 `gorm:"primaryKey"`
	UserID      uuid.UUID
	RecipeID    int64
	RecipeTitle string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
