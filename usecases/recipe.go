package usecases

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	cs "github.com/satryanararya/go-chefbot/drivers/cloudinary"
	client "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	dto "github.com/satryanararya/go-chefbot/dto/recipe"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type RecipeUseCase interface {
	SearchRecipe(c echo.Context, name string) (client.SearchRecipeResponse, error)
	GetRecipeInformation(c echo.Context, recipeID int) (client.RecipeInformation, error)
	CreateRecipe(c echo.Context, id int64, req *dto.CreateRecipeRequest) (*dto.CreateRecipeResponse, error)
	UpdateRecipe(c echo.Context, id int64, req *dto.UpdateRecipeRequest) error
	GetRecipe(c echo.Context, id int64) (*dto.GetRecipeResponse, error)
	DeleteRecipe(c echo.Context, id int64) error
}

type recipeUseCase struct {
	client            client.RecipeClient
	recipeRepo        repositories.RecipeRepository
	cloudinaryService cs.CloudinaryService
}

func NewRecipeUseCase(c client.RecipeClient, rr repositories.RecipeRepository, cs cs.CloudinaryService) *recipeUseCase {
	return &recipeUseCase{
		client:            c,
		recipeRepo:        rr,
		cloudinaryService: cs,
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

func (ruc *recipeUseCase) CreateRecipe(c echo.Context, id int64, req *dto.CreateRecipeRequest) (*dto.CreateRecipeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	formHeader, err := c.FormFile("image")
	if err != nil {
		return nil, err
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return nil, err
	}


	imageURL, err := ruc.cloudinaryService.UploadImage(ctx, formFile)
	if err != nil {
		return nil, err
	}

	recipe := &entities.Recipe{
		UserID:          id,
		Title:           req.Title,
		Image:           imageURL,
		SourceName:      req.SourceName,
		CookingMinutes:  req.CookingMinutes,
		PricePerServing: req.PricePerServing,
		ReadyInMinutes:  req.ReadyInMinutes,
		Servings:        req.Servings,
		HealthScore:     req.HealthScore,
		Diets:           req.Diets,
		IsSustainable:   req.IsSustainable,
		Instruction:     req.Instruction,
	}

	err = ruc.recipeRepo.CreateRecipe(ctx, recipe)
	if err != nil {
		return nil, err
	}

	return &dto.CreateRecipeResponse{
		ID:    recipe.ID,
		Title: recipe.Title,
		Image: recipe.Image,
	}, nil
}

func (ruc *recipeUseCase) UpdateRecipe(c echo.Context, id int64, req *dto.UpdateRecipeRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	recipe, err := ruc.recipeRepo.GetRecipe(ctx, id)
	if err != nil {
		return err
	}

	formHeader, err := c.FormFile("image")
	if err != nil {
		return err
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return  err
	}

	imageURL, err := ruc.cloudinaryService.UploadImage(ctx, formFile)
	if err != nil {
		return err
	}

	// Update the recipe fields
	recipe.Title = req.Title
	recipe.Image = imageURL // Use the URL of the uploaded image
	recipe.SourceName = req.SourceName
	recipe.CookingMinutes = req.CookingMinutes
	recipe.PricePerServing = req.PricePerServing
	recipe.ReadyInMinutes = req.ReadyInMinutes
	recipe.Servings = req.Servings
	recipe.HealthScore = req.HealthScore
	recipe.Diets = req.Diets
	recipe.IsSustainable = req.IsSustainable
	recipe.Instruction = req.Instruction

	// Update the extended ingredients
	// recipe.ExtendedIngredients = make([]entities.ExtendedIngredient, len(req.ExtendedIngredients))
	// for i, ingredientDTO := range req.ExtendedIngredients {
	//     recipe.ExtendedIngredients[i] = entities.ExtendedIngredient{
	//         Name:   ingredientDTO.Name,
	//         Amount: ingredientDTO.Amount,
	//         Unit:   ingredientDTO.Unit,
	//     }
	// }

	// Update the recipe in the database
	return ruc.recipeRepo.UpdateRecipe(ctx, recipe)
}

func (ruc *recipeUseCase) GetRecipe(c echo.Context, id int64) (*dto.GetRecipeResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	// Fetch the recipe from the database
	recipe, err := ruc.recipeRepo.GetRecipe(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &dto.GetRecipeResponse{
		ID:              recipe.ID,
		UserID:          recipe.UserID,
		Title:           recipe.Title,
		Image:           recipe.Image,
		SourceName:      recipe.SourceName,
		CookingMinutes:  recipe.CookingMinutes,
		PricePerServing: recipe.PricePerServing,
		ReadyInMinutes:  recipe.ReadyInMinutes,
		Servings:        recipe.Servings,
		HealthScore:     recipe.HealthScore,
		Diets:           recipe.Diets,
		IsSustainable:   recipe.IsSustainable,
		Instruction:     recipe.Instruction,
	}

	// Convert the extended ingredients
	// res.ExtendedIngredients = make([]dto.ExtendedIngredientDTO, len(recipe.ExtendedIngredients))
	// for i, ingredient := range recipe.ExtendedIngredients {
	// 	res.ExtendedIngredients[i] = dto.ExtendedIngredientDTO{
	// 		Name:   ingredient.Name,
	// 		Amount: ingredient.Amount,
	// 		Unit:   ingredient.Unit,
	// 	}
	// }

	return res, nil
}

func (ruc *recipeUseCase) DeleteRecipe(c echo.Context, id int64) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return ruc.recipeRepo.DeleteRecipe(ctx, id)
}
