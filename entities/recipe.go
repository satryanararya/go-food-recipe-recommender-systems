package entities

type Recipe struct {
	ID                  int64  `gorm:"primaryKey"`
	UserID              int64  `gorm:"index"`
	Title               string `gorm:"not null"`
	Image               string
	SourceName          string
	CookingMinutes      int64
	ExtendedIngredients []ExtendedIngredient
	PricePerServing     float64
	ReadyInMinutes      int
	Servings            int
	HealthScore         float64
	Diets               []string
	IsSustainable       bool
	Instruction         string
}

type ExtendedIngredient struct {
	ID     int64  `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	Amount float64
	Unit   string
}
