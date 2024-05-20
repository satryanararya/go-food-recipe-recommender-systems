package recipe

import "mime/multipart"

type RecipeRequest struct {
	Title           string         `json:"title" form:"title" validate:"required,min=2,max=50"`
	SourceName      *string         `json:"source_name" form:"source_name"`
	CookingMinutes  int64          `json:"cooking_minutes" form:"cooking_minutes"`
	Ingredient      []Ingredient   `json:"ingredient" form:"ingredient"`
	PricePerServing *float64        `json:"price_per_serving" form:"price_per_serving"`
	ReadyInMinutes  int            `json:"ready_in_minutes" form:"ready_in_minutes"`
	Servings        int            `json:"servings" form:"servings"`
	Diets           []string       `json:"diets" form:"diets"`
	IsSustainable   bool           `json:"is_sustainable" form:"is_sustainable"`
	Instruction     string         `json:"instruction" form:"instruction" validate:"required,min=2,max=500"`
}

type RecipeResponse struct {
	Title string `json:"title"`
	Image *string `json:"image"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type RecipeImageRequest struct {
	Image *multipart.FileHeader `json:"image" form:"image"`
}