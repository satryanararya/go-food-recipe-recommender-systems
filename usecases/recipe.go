package usecases

import (
    "context"

    "github.com/labstack/echo/v4"
    client "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
)

type RecipeUseCase interface {
    SearchRecipe(c echo.Context, name string) (client.SearchRecipeResponse, error)
    GetRecipeInformation(c echo.Context, recipeID int) (client.RecipeInformation, error)
}

type recipeUseCase struct {
    client client.RecipeClient
}

func NewRecipeUseCase(c client.RecipeClient) *recipeUseCase {
    return &recipeUseCase{
        client: c,
    }
}

func (ruc *recipeUseCase) SearchRecipe(c echo.Context, name string) (client.SearchRecipeResponse, error) {
    ctx, cancel := context.WithCancel(c.Request().Context())
    defer cancel()

    searchResponse, err := ruc.client.SearchRecipe(ctx, name)
    if err != nil {
        return client.SearchRecipeResponse{}, err
    }

    var recipes []client.Recipe
    for _, response := range searchResponse.Results {
        recipes = append(recipes, client.Recipe{
            ID:    response.ID,
            Title: response.Title,
            Image: response.Image,
        })
    }

    return client.SearchRecipeResponse{
        Results:      recipes,
        Offset:       searchResponse.Offset,
        Number:       searchResponse.Number,
        TotalResults: searchResponse.TotalResults,
    }, nil
}

func (ruc *recipeUseCase) GetRecipeInformation(c echo.Context, recipeID int) (client.RecipeInformation, error) {
    ctx, cancel := context.WithCancel(c.Request().Context())
    defer cancel()

    recipeInformationResponse, err := ruc.client.GetRecipeInformation(ctx, recipeID)
    if err != nil {
        return client.RecipeInformation{}, err
    }

    recipeInformation := client.RecipeInformation{
        Title:               recipeInformationResponse.Title,
        Image:               recipeInformationResponse.Image,
        SourceName:          recipeInformationResponse.SourceName,
        CookingMinutes:      recipeInformationResponse.CookingMinutes,
        ExtendedIngredients: make([]client.ExtendedIngredient, len(recipeInformationResponse.ExtendedIngredients)),
        PricePerServing:     recipeInformationResponse.PricePerServing,
        ReadyInMinutes:      recipeInformationResponse.ReadyInMinutes,
        Servings:            recipeInformationResponse.Servings,
        HealthScore:         recipeInformationResponse.HealthScore,
        Diets:               recipeInformationResponse.Diets,
        IsSustainable:       recipeInformationResponse.IsSustainable,
        Instruction:         recipeInformationResponse.Instruction,
    }

    for i, ingredient := range recipeInformationResponse.ExtendedIngredients {
        recipeInformation.ExtendedIngredients[i] = client.ExtendedIngredient{
            Name:   ingredient.Name,
            Amount: ingredient.Amount,
            Unit:   ingredient.Unit,
        }
    }

    return recipeInformation, nil
}