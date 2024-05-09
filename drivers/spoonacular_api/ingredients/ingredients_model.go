package ingredients

type Ingredient struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Image     string       `json:"image"`
	Category  string       `json:"categoryPath"`
	Cost      *[]Cost      `json:"estimatedCost"`
	Nutrition *[]Nutrition `json:"nutrition"`
}

type Cost struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Nutrition struct {
	Nutrients        *[]Nutrients        `json:"nutrients"`
	Calories         *[]Calories         `json:"caloricBreakdown"`
	WeightPerServing *[]WeightPerServing `json:"weightPerServing"`
}

type Nutrients struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type Calories struct {
	Protein float64 `json:"percentProtein"`
	Fat     float64 `json:"percentFat"`
	Carbs   float64 `json:"percentCarbs"`
}

type WeightPerServing struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}