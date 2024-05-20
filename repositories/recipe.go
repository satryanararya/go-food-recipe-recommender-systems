package repositories

import (
	"context"

	"github.com/google/uuid"
	dto_p "github.com/satryanararya/go-chefbot/dto"
	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	CreateRecipe(ctx context.Context, recipe *entities.Recipe) error
	GetRecipe(ctx context.Context, recipeID int) (*entities.Recipe, error)
	GetUserRecipes(ctx context.Context, id uuid.UUID, p *dto_p.PaginationRequest) ([]entities.Recipe, int64, error)
	UpdateRecipe(ctx context.Context, recipe *entities.Recipe) error
    DeleteRecipe(ctx context.Context, recipe *entities.Recipe) error
}

type recipeRepository struct {
	DB *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		DB: db,
	}
}

func (r *recipeRepository) CreateRecipe(ctx context.Context, recipe *entities.Recipe) error {
    tx := r.DB.Begin()

    // Save the recipe in the database
    if err := tx.WithContext(ctx).Omit("RecipeIngredients").Create(recipe).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Save each recipe ingredient in the database
    for _, ingredient := range recipe.RecipeIngredients {
        ingredient.RecipeID = recipe.ID
        if err := tx.WithContext(ctx).Omit("ID").Create(&ingredient).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    tx.Commit()

    return nil
}

func (r *recipeRepository) GetRecipe(ctx context.Context, recipeID int) (*entities.Recipe, error) {
    var recipe entities.Recipe
    if err := r.DB.WithContext(ctx).Preload("RecipeIngredients").First(&recipe, recipeID).Error; err != nil {
        return nil, err
    }
    return &recipe, nil
}

func (r *recipeRepository) GetUserRecipes(ctx context.Context, id uuid.UUID, p *dto_p.PaginationRequest) ([]entities.Recipe, int64, error) {
    if err := ctx.Err(); err != nil {
        return nil, 0, err
    }

    offset := (p.Page - 1) * p.Limit

    var recipes []entities.Recipe
    result := r.DB.Preload("RecipeIngredients").Preload("RatingReview").Where("user_id = ?", id).Limit(p.Limit).Offset(offset).Find(&recipes)
    if result.Error != nil {
        return nil, 0, result.Error
    }

    var total int64
    r.DB.Model(&entities.Recipe{}).Where("user_id = ?", id).Count(&total)

    return recipes, total, nil
}

func (r *recipeRepository) UpdateRecipe(ctx context.Context, recipe *entities.Recipe) error {
    if err := r.DB.WithContext(ctx).Save(recipe).Error; err != nil {
        return err
    }

    return nil
}

func (r *recipeRepository) DeleteRecipe(ctx context.Context, recipe *entities.Recipe) error {
    // Begin a new transaction
    tx := r.DB.Begin()

    // Delete the associated rows in the recipe_ingredients table
    if err := tx.WithContext(ctx).Where("recipe_id = ?", recipe.ID).Delete(&entities.RecipeIngredient{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Delete the recipe
    if err := tx.WithContext(ctx).Delete(recipe).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Commit the transaction
    tx.Commit()

    return nil
}