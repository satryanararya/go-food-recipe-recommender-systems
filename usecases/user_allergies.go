package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/satryanararya/go-chefbot/drivers/spoonacular_api/ingredients"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type UserAllergiesUseCase interface {
	GetIngredientInfo(ctx context.Context, userID int64, dto *dto.UserAllergiesRequest) (*entities.Ingredient, error)
}

type userAllergiesUseCase struct {
	client            ingredients.IngredientClient
	ingredientRepo    repositories.IngredientRepository
	userAllergiesRepo repositories.UserAllergyRepository
}

func NewUserAllergiesUseCase(c ingredients.IngredientClient, ir repositories.IngredientRepository, uar repositories.UserAllergyRepository) *userAllergiesUseCase {
	return &userAllergiesUseCase{
		client:            c,
		ingredientRepo:    ir,
		userAllergiesRepo: uar,
	}
}

func (uac *userAllergiesUseCase) GetIngredientInfo(ctx context.Context, userID int64, dto *dto.UserAllergiesRequest) (*entities.Ingredient, error) {
	exists, err := uac.ingredientRepo.ExistsByName(ctx, dto.IngredientName)
	if err != nil {
		return nil, err
	}

	var ingredient *entities.Ingredient
	if exists {
		ingredient, err = uac.ingredientRepo.GetByName(ctx, dto.IngredientName)
	} else {
		ingredientID, err := uac.client.SearchIngredient(ctx, dto)
		if err != nil {
			return nil, err
		}

		ingredientInfo, err := uac.client.GetIngredient(ctx, ingredientID)
		if err != nil {
			return nil, err
		}

		// Convert drivers/ingredients.Ingredient to entities.Ingredient
		ingredient = &entities.Ingredient{
			ID:   int64(ingredientInfo.ID),
			Name: ingredientInfo.Name,
		}

		err = uac.ingredientRepo.Save(ctx, ingredient)
		if err != nil {
			log.Printf("Failed to save ingredient: %v", err)
			return nil, fmt.Errorf("failed to save ingredient: %w", err)
		}
	}

	err = uac.userAllergiesRepo.Save(ctx, userID, ingredient.ID)
	if err != nil {
		log.Printf("Failed to save user allergy: %v", err)
		return nil, fmt.Errorf("failed to save user allergy: %w", err)
	}

	return ingredient, nil
}
