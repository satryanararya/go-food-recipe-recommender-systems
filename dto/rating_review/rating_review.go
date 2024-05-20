package ratingreview

import "github.com/google/uuid"

type RatingReviewRequest struct {
	Rating float64 `json:"rating"`
	Review string  `json:"review"`
}

type RatingReviewResponse struct {
	UserID   uuid.UUID   `json:"user_id"`
	RecipeID int64   `json:"recipe_id"`
	Rating   float64 `json:"rating"`
	Review   string  `json:"review"`
}
