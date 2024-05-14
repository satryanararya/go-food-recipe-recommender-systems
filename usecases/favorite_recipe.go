package usecases

import (
	"context"

	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/favorite_recipe"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type FavoriteRecipeUsecase interface {
	AddToFavorites(c echo.Context, userID int64, recipeID int64) error
	RemoveFromFavorites(c echo.Context, userID int64, recipeID int64) error
	FindFavoritesByUserID(c echo.Context, userID int64) ([]dto.FavoriteRecipeResponse, error)
}

type favoriteRecipeUsecase struct {
	favoriteRecipeRepo repositories.FavoriteRecipeRepository
}

func NewFavoriteRecipeUsecase(fr repositories.FavoriteRecipeRepository) *favoriteRecipeUsecase {
	return &favoriteRecipeUsecase{
		favoriteRecipeRepo: fr,
	}
}

func (fruc *favoriteRecipeUsecase) AddToFavorites(c echo.Context, userID int64, recipeID int64) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	favoriteRecipe := &entities.FavoriteRecipe{
		UserID:   userID,
		RecipeID: recipeID,
	}
	return fruc.favoriteRecipeRepo.AddToFavorites(ctx, favoriteRecipe)
}

func (fruc *favoriteRecipeUsecase) RemoveFromFavorites(c echo.Context, userID int64, recipeID int64) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return fruc.favoriteRecipeRepo.RemoveFromFavorites(ctx, userID, recipeID)
}

func (fruc *favoriteRecipeUsecase) FindFavoritesByUserID(c echo.Context, userID int64) ([]dto.FavoriteRecipeResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	favoriteRecipes, err := fruc.favoriteRecipeRepo.FindFavoritesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.FavoriteRecipeResponse
	for _, fr := range favoriteRecipes {
		responses = append(responses, dto.FavoriteRecipeResponse{
			UserID:   fr.UserID,
			RecipeID: fr.RecipeID,
		})
	}

	return responses, nil
}