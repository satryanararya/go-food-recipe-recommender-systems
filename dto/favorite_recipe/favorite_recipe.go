package favoriterecipe

type FavoriteRecipeResponse struct {
	UserID   int64   `json:"user_id"`
	RecipeID int64   `json:"recipe_id"`
}
