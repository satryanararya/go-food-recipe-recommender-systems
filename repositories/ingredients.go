package repositories

import (
    "context"

    "github.com/satryanararya/go-chefbot/entities"
    "gorm.io/gorm"
)

type IngredientRepository interface {
    Save(ctx context.Context, ingredient *entities.Ingredient) error
    Get(ctx context.Context, id int) (*entities.Ingredient, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
    GetByName(ctx context.Context, name string) (*entities.Ingredient, error)
}

type ingredientRepository struct {
    DB *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IngredientRepository {
    return &ingredientRepository{
        DB: db,
    }
}

func (r *ingredientRepository) Save(ctx context.Context, ingredient *entities.Ingredient) error {
    result := r.DB.WithContext(ctx).Save(ingredient)
    return result.Error
}

func (r *ingredientRepository) Get(ctx context.Context, id int) (*entities.Ingredient, error) {
    var ingredient entities.Ingredient
    result := r.DB.WithContext(ctx).First(&ingredient, id)
    return &ingredient, result.Error
}

func (r *ingredientRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
    var count int64
    result := r.DB.WithContext(ctx).Model(&entities.Ingredient{}).Where("name = ?", name).Count(&count)
    return count > 0, result.Error
}

func (r *ingredientRepository) GetByName(ctx context.Context, name string) (*entities.Ingredient, error) {
    var ingredient entities.Ingredient
    result := r.DB.WithContext(ctx).Where("name = ?", name).First(&ingredient)
    return &ingredient, result.Error
}