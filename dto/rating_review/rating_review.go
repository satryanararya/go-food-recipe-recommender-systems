package ratingreview

type RatingReviewRequest struct {
	Rating float64 `json:"rating"`
	Review string  `json:"review"`
}

type RatingReviewResponse struct {
	UserID   int64   `json:"user_id"`
	RecipeID int64   `json:"recipe_id"`
	Rating   float64 `json:"rating"`
	Review   string  `json:"review"`
}
