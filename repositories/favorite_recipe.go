package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type FavoriteRecipeRepository interface {
	AddToFavorites(ctx context.Context, favoriteRecipe *entities.FavoriteRecipe) error
	RemoveFromFavorites(ctx context.Context, userID uuid.UUID, recipeID int64) error
	FindFavoritesByUserID(ctx context.Context, userID uuid.UUID) ([]entities.FavoriteRecipe, error)
}

type favoriteRecipeRepository struct {
	DB *gorm.DB
}

func NewFavoriteRecipeRepository(db *gorm.DB) *favoriteRecipeRepository {
	return &favoriteRecipeRepository{
		DB: db,
	}
}

func (r *favoriteRecipeRepository) AddToFavorites(ctx context.Context, favoriteRecipe *entities.FavoriteRecipe) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.DB.Create(favoriteRecipe).Error
}

func (r *favoriteRecipeRepository) RemoveFromFavorites(ctx context.Context, userID uuid.UUID, recipeID int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.DB.Where("user_id = ? AND recipe_id = ?", userID, recipeID).Delete(&entities.FavoriteRecipe{}).Error
}

func (r *favoriteRecipeRepository) FindFavoritesByUserID(ctx context.Context, userID uuid.UUID) ([]entities.FavoriteRecipe, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	var favoriteRecipes []entities.FavoriteRecipe
	if err := r.DB.Where("user_id = ?", userID).Find(&favoriteRecipes).Error; err != nil {
		return nil, err
	}

	return favoriteRecipes, nil
}
