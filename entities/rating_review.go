package entities

import (
	"time"

	"github.com/google/uuid"
)

type RatingReview struct {
	ID        int64 `gorm:"primaryKey" json:"-"`
	UserID    uuid.UUID `json:"user_id"`
	RecipeID  int64 `json:"-"`
	Rating    float64 `json:"rating"`
	Review    string `json:"review"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}
