package entities

type Ingredient struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
