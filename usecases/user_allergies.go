package usecases

import (
	"context"
	"fmt"
	"strings"

	// "strings"

	"github.com/google/uuid"
	"github.com/satryanararya/go-chefbot/drivers/spoonacular_api/ingredients"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type UserAllergiesUseCase interface {
	AddAllergies(ctx context.Context, userID uuid.UUID, dto *dto.UserAllergiesRequest) (*entities.UserAllergies, error)
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

func (uac *userAllergiesUseCase) AddAllergies(ctx context.Context, userID uuid.UUID, dto *dto.UserAllergiesRequest) (*entities.UserAllergies, error) {
	ingredient, err := uac.getOrCreateIngredient(ctx, dto.IngredientName)
	if err != nil {
		return nil, err
	}

	userAllergies := &entities.UserAllergies{
		UserID:     userID,
		IngredientID: ingredient.ID,
		Ingredient: *ingredient,
	}
	fmt.Println("User allergies: ", userAllergies)

	err = uac.userAllergiesRepo.Save(ctx, userAllergies)
	if err != nil {
		fmt.Println("Error saving user allergies: ", err)
		return nil, err
	}
	

	return userAllergies, nil
}

func (uac *userAllergiesUseCase) getOrCreateIngredient(ctx context.Context, name string) (*entities.Ingredient, error) {
	exists, err := uac.ingredientRepo.ExistsByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var ingredient *entities.Ingredient
	if exists {
		ingredient, err = uac.ingredientRepo.GetByName(ctx, name)
		if err != nil {
			return nil, err
		}
	} else {
		ingredientID, err := uac.client.SearchIngredient(ctx, &dto.UserAllergiesRequest{IngredientName: name})
		if err != nil {
			return nil, err
		}

		ingredientInfo, err := uac.client.GetIngredientInfo(ctx, ingredientID)
		if err != nil {
			fmt.Println("Error getting ingredient info: ", err)
			return nil, err
		}

		// Extract the necessary information from the Ingredient model
		category := strings.Join(ingredientInfo.Category, ", ")
		estimatedCost := fmt.Sprintf("%.2f %s", ingredientInfo.Cost.Value, ingredientInfo.Cost.Unit)
		weightPerServing := fmt.Sprintf("%.2f %s", ingredientInfo.Nutrition.WeightPerServing.Amount, ingredientInfo.Nutrition.WeightPerServing.Unit)

		ingredient = &entities.Ingredient{
			ID:       int64(ingredientInfo.ID),
			Name:     ingredientInfo.Name,
			Category: category,
			IngredientDetails: entities.IngredientDetails{
				EstimatedCost:    estimatedCost,
				WeightPerServing: weightPerServing,
				Protein:          ingredientInfo.Nutrition.Calories.Protein,
				Fat:              ingredientInfo.Nutrition.Calories.Fat,
				Carbs:            ingredientInfo.Nutrition.Calories.Carbs,
			},
		}

		err = uac.ingredientRepo.Save(ctx, ingredient)
		if err != nil {
			return nil, err
		}
	}

	return ingredient, nil
}
