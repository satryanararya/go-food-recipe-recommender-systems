package recipe

import "mime/multipart"

type CreateRecipeRequest struct {
	Title               string                  `json:"title" form:"title" validate:"required,min=2,max=50"`
	Image               multipart.File   `json:"image" form:"image"`
	SourceName          string                  `json:"source_name" form:"source_name"`
	CookingMinutes      int64                   `json:"cooking_minutes" form:"cooking_minutes"`
	// ExtendedIngredients []ExtendedIngredientDTO `json:"extended_ingredients" form:"extended_ingredients"`
	PricePerServing     float64                 `json:"price_per_serving" form:"price_per_serving"`
	ReadyInMinutes      int                     `json:"ready_in_minutes" form:"ready_in_minutes"`
	Servings            int                     `json:"servings" form:"servings"`
	HealthScore         float64                 `json:"health_score" form:"health_score"`
	Diets               string                  `json:"diets" form:"diets"`
	IsSustainable       bool                    `json:"is_sustainable" form:"is_sustainable"`
	Instruction         string                  `json:"instruction" form:"instruction" validate:"required,min=2,max=500"`
}

type CreateRecipeResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type UpdateRecipeRequest struct {
	Title               string                  `json:"title" form:"title" validate:"required,min=2,max=50"`
	Image               *multipart.FileHeader   `json:"image" form:"image"`
	SourceName          string                  `json:"source_name" form:"source_name"`
	CookingMinutes      int64                   `json:"cooking_minutes" form:"cooking_minutes"`
	// ExtendedIngredients []ExtendedIngredientDTO `json:"extended_ingredients" form:"extended_ingredients"`
	PricePerServing     float64                 `json:"price_per_serving" form:"price_per_serving"`
	ReadyInMinutes      int                     `json:"ready_in_minutes" form:"ready_in_minutes"`
	Servings            int                     `json:"servings" form:"servings"`
	HealthScore         float64                 `json:"health_score" form:"health_score"`
	Diets               string                  `json:"diets" form:"diets"`
	IsSustainable       bool                    `json:"is_sustainable" form:"is_sustainable"`
	Instruction         string                  `json:"instruction" form:"instruction" validate:"required,min=2,max=500"`
}

type GetRecipeResponse struct {
	ID                  int64                   `json:"id"`
	UserID              int64                   `json:"user_id"`
	Title               string                  `json:"title"`
	Image               string                  `json:"image"`
	SourceName          string                  `json:"source_name"`
	CookingMinutes      int64                   `json:"cooking_minutes"`
	// ExtendedIngredients []ExtendedIngredientDTO `json:"extended_ingredients"`
	PricePerServing     float64                 `json:"price_per_serving"`
	ReadyInMinutes      int                     `json:"ready_in_minutes"`
	Servings            int                     `json:"servings"`
	HealthScore         float64                 `json:"health_score"`
	Diets               string                  `json:"diets"`
	IsSustainable       bool                    `json:"is_sustainable"`
	Instruction         string                  `json:"instruction"`
}

// type ExtendedIngredientDTO struct {
// 	Name   string  `json:"name"`
// 	Amount float64 `json:"amount"`
// 	Unit   string  `json:"unit"`
// }
