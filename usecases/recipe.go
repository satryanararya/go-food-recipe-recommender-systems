package usecases

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	cs "github.com/satryanararya/go-chefbot/drivers/cloudinary"
	client "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	dto_p "github.com/satryanararya/go-chefbot/dto"
	dto "github.com/satryanararya/go-chefbot/dto/recipe"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
	err_util "github.com/satryanararya/go-chefbot/utils/error"
)

type RecipeUseCase interface {
	// External
	SearchRecipe(c echo.Context, name string) (client.SearchRecipeResponse, error)
	GetRecipeInformation(c echo.Context, recipeID int) (client.RecipeInformation, error)

	CreateRecipe(c echo.Context, id uuid.UUID, req *dto.RecipeRequest) (*dto.RecipeResponse, error)
	UploadRecipeImage(c echo.Context, id uuid.UUID, recipeID int, req *dto.RecipeImageRequest) error
	GetUserRecipes(c echo.Context, id uuid.UUID, p *dto_p.PaginationRequest) ([]entities.Recipe, *dto_p.PaginationMetadata, *dto_p.Link, error)
	UpdateRecipe(c echo.Context, id uuid.UUID, recipeID int, req *dto.RecipeRequest) (*dto.RecipeResponse, error)
	DeleteRecipe(c echo.Context, id uuid.UUID, recipeID int) error
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

func (ruc *recipeUseCase) CreateRecipe(c echo.Context, id uuid.UUID, req *dto.RecipeRequest) (*dto.RecipeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	diets := strings.Join(req.Diets, ", ")

	ingredients := make([]entities.RecipeIngredient, len(req.Ingredient))
	for i, ingredientReq := range req.Ingredient {
		ingredients[i] = entities.RecipeIngredient{
			Name:     ingredientReq.Name,
			Quantity: ingredientReq.Quantity,
			Unit:     ingredientReq.Unit,
		}
	}

	recipe := &entities.Recipe{
		UserID:            id,
		Title:             req.Title,
		SourceName:        *req.SourceName,
		CookingMinutes:    req.CookingMinutes,
		PricePerServing:   *req.PricePerServing,
		ReadyInMinutes:    req.ReadyInMinutes,
		Servings:          req.Servings,
		Diets:             diets,
		IsSustainable:     req.IsSustainable,
		RecipeIngredients: ingredients,
		Instruction:       req.Instruction,
	}

	err := ruc.recipeRepo.CreateRecipe(ctx, recipe)
	if err != nil {
		return nil, err
	}
	return &dto.RecipeResponse{
		Title: recipe.Title,
		Image: &recipe.Image,
	}, nil
}

func (ruc *recipeUseCase) UploadRecipeImage(c echo.Context, id uuid.UUID, recipeID int, req *dto.RecipeImageRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	formHeader, err := c.FormFile("image")
	if err != nil {
		return err
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return err
	}

	imageURL, err := ruc.cloudinaryService.UploadImage(ctx, formFile)
	if err != nil {
		return err
	}

	recipe, err := ruc.recipeRepo.GetRecipe(ctx, recipeID)
	if err != nil {
		return err
	}

	// Check if the user is the owner of the recipe
	if recipe.UserID != id {
		return err
	}

	recipe.Image = imageURL

	err = ruc.recipeRepo.UpdateRecipe(ctx, recipe)
	if err != nil {
		return err
	}

	return nil
}

func (ruc *recipeUseCase) GetUserRecipes(c echo.Context, id uuid.UUID, p *dto_p.PaginationRequest) ([]entities.Recipe, *dto_p.PaginationMetadata, *dto_p.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	baseURL := fmt.Sprintf(
		"%s?limit=%d&page=",
		c.Request().URL.Path,
		p.Limit,
	)
	var (
		next = baseURL + strconv.Itoa(p.Page+1)
		prev = baseURL + strconv.Itoa(p.Page-1)
	)
	recipes, totalData, err := ruc.recipeRepo.GetUserRecipes(ctx, id, p)
	if err != nil {
		return nil, nil, nil, err
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(p.Limit)))
	meta := &dto_p.PaginationMetadata{
		CurrentPage: p.Page,
		TotalPage:   totalPage,
		TotalData:   totalData,
	}

	if p.Page > totalPage {
		return nil, nil, nil, err_util.ErrPageNotFound
	}

	if p.Page == 1 {
		prev = ""
	}

	if p.Page == totalPage {
		next = ""
	}

	link := &dto_p.Link{
		Next: next,
		Prev: prev,
	}

	return recipes, meta, link, nil
}

func (ruc *recipeUseCase) UpdateRecipe(c echo.Context, id uuid.UUID, recipeID int, req *dto.RecipeRequest) (*dto.RecipeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	recipe, err := ruc.recipeRepo.GetRecipe(ctx, recipeID)
	if err != nil {
		return nil, err
	}
	if recipe.UserID != id {
		return nil, err
	}

	recipe.Title = req.Title
	recipe.SourceName = *req.SourceName
	recipe.CookingMinutes = req.CookingMinutes
	recipe.PricePerServing = *req.PricePerServing
	recipe.ReadyInMinutes = req.ReadyInMinutes
	recipe.Servings = req.Servings
	recipe.Diets = strings.Join(req.Diets, ", ")
	recipe.IsSustainable = req.IsSustainable
	recipe.Instruction = req.Instruction

	ingredients := make([]entities.RecipeIngredient, len(req.Ingredient))
	for i, ingredientReq := range req.Ingredient {
		ingredients[i] = entities.RecipeIngredient{
			Name:     ingredientReq.Name,
			Quantity: ingredientReq.Quantity,
			Unit:     ingredientReq.Unit,
		}
	}
	recipe.RecipeIngredients = ingredients

	err = ruc.recipeRepo.UpdateRecipe(ctx, recipe)
	if err != nil {
		return nil, err
	}

	return &dto.RecipeResponse{
		Title: recipe.Title,
		Image: &recipe.Image,
	}, nil
}

func (ruc *recipeUseCase) DeleteRecipe(c echo.Context, id uuid.UUID, recipeID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	recipe, err := ruc.recipeRepo.GetRecipe(ctx, recipeID)
	if err != nil {
		return err
	}
	if recipe.UserID != id {
		return err
	}

	err = ruc.recipeRepo.DeleteRecipe(ctx, recipe)
	if err != nil {
		return err
	}

	return nil
}