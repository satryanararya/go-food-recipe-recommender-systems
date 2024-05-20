package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/favorite_recipe"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type FavoriteRecipeUsecase interface {
	AddToFavorites(c echo.Context, userID uuid.UUID, recipeID int64) error
	RemoveFromFavorites(c echo.Context, userID uuid.UUID, recipeID int64) error
	FindFavoritesByUserID(c echo.Context, userID uuid.UUID) ([]dto.FavoriteRecipeResponse, error)
}

type favoriteRecipeUsecase struct {
	favoriteRecipeRepo repositories.FavoriteRecipeRepository
	recipeRepo 	   repositories.RecipeRepository
}

func NewFavoriteRecipeUsecase(fr repositories.FavoriteRecipeRepository, rr repositories.RecipeRepository) *favoriteRecipeUsecase {
	return &favoriteRecipeUsecase{
		favoriteRecipeRepo: fr,
		recipeRepo:         rr,
	}
}

func (fruc *favoriteRecipeUsecase) AddToFavorites(c echo.Context, userID uuid.UUID, recipeID int64) error {
    ctx, cancel := context.WithCancel(c.Request().Context())
    defer cancel()

    // Retrieve the recipe from the database
    recipe, err := fruc.recipeRepo.GetRecipe(ctx, int(recipeID))
    if err != nil {
        return err
    }

    favoriteRecipe := &entities.FavoriteRecipe{
        UserID:     userID,
        RecipeID:   recipeID,
        RecipeTitle: recipe.Title, // Set the RecipeTitle field
    }
    return fruc.favoriteRecipeRepo.AddToFavorites(ctx, favoriteRecipe)
}

func (fruc *favoriteRecipeUsecase) RemoveFromFavorites(c echo.Context, userID uuid.UUID, recipeID int64) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return fruc.favoriteRecipeRepo.RemoveFromFavorites(ctx, userID, recipeID)
}

func (fruc *favoriteRecipeUsecase) FindFavoritesByUserID(c echo.Context, userID uuid.UUID) ([]dto.FavoriteRecipeResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	favoriteRecipes, err := fruc.favoriteRecipeRepo.FindFavoritesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.FavoriteRecipeResponse
	for _, fr := range favoriteRecipes {
		responses = append(responses, dto.FavoriteRecipeResponse{
			RecipeID: fr.RecipeID,
			RecipeTitle: fr.RecipeTitle,
		})
	}

	return responses, nil
}