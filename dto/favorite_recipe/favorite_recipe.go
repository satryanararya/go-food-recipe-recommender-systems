package favoriterecipe

type FavoriteRecipeResponse struct {
	RecipeID int64   `json:"recipe_id"`
	RecipeTitle string `json:"recipe_title"`
}
