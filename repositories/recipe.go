package repositories

import (
	"context"

	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	CreateRecipe(ctx context.Context, recipe *entities.Recipe) error
	UpdateRecipe(ctx context.Context, recipe *entities.Recipe) error
	GetRecipe(ctx context.Context, id int64) (*entities.Recipe, error)
	DeleteRecipe(ctx context.Context, id int64) error
}

type recipeRepository struct {
	DB                *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		DB:                db,
	}
}

func (r *recipeRepository) CreateRecipe(ctx context.Context, recipe *entities.Recipe) error {
    if err := r.DB.WithContext(ctx).Create(recipe).Error; err != nil {
        return err
    }

    return nil
}

func (r *recipeRepository) UpdateRecipe(ctx context.Context, recipe *entities.Recipe) error {
    if err := r.DB.WithContext(ctx).Save(recipe).Error; err != nil {
        return err
    }

    return nil
}

func (r *recipeRepository) GetRecipe(ctx context.Context, id int64) (*entities.Recipe, error) {
    var recipe entities.Recipe

    if err := r.DB.WithContext(ctx).First(&recipe, id).Error; err != nil {
        return nil, err
    }

    return &recipe, nil
}

func (r *recipeRepository) DeleteRecipe(ctx context.Context, id int64) error {
    if err := r.DB.WithContext(ctx).Delete(&entities.Recipe{}, id).Error; err != nil {
        return err
    }

    return nil
}