package recipe

type Recipe struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type RecipeInformation struct {
	Title               string               `json:"title"`
	Image               string               `json:"image"`
	SourceName          string               `json:"sourceName"`
	CookingMinutes      int64                `json:"cookingMinutes"`
	ExtendedIngredients []ExtendedIngredient `json:"extendedIngredients"`
	PricePerServing     float64              `json:"pricePerServing"`
	ReadyInMinutes      int64                `json:"readyInMinutes"`
	Servings            int64                `json:"servings"`
	HealthScore         float64              `json:"healthScore"`
	Diets               []string             `json:"diets"`
	IsSustainable       bool                 `json:"sustainable"`
	Instruction         string               `json:"instructions"`
}

type ExtendedIngredient struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type SearchRecipeResponse struct {
	Results      []Recipe `json:"results"`
	Offset       int64     `json:"offset"`
	Number       int64     `json:"number"`
	TotalResults int64     `json:"totalResults"`
}
