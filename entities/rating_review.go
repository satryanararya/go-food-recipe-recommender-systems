package entities

import "time"

type RatingReview struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	RecipeID  int64
	Rating    float64
	Review    string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
