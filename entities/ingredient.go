package entities

type Ingredient struct {
	ID       int64  `gorm:"primaryKey" json:"-"` 
	Name     string `gorm:"not null" json:"name"`
	Category string `gorm:"not null" json:"category"`
	IngredientDetails `gorm:"embedded"`
}

type IngredientDetails struct {
	Protein          float64 `gorm:"not null" json:"protein"`
	Fat              float64 `gorm:"not null" json:"fat"`
	Carbs            float64 `gorm:"not null" json:"carbs"`
	WeightPerServing string  `gorm:"not null" json:"weightPerServing"`
	EstimatedCost    string  `gorm:"not null" json:"estimatedCost"`
}
