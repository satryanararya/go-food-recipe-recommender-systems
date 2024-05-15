package recipe

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	err_util "github.com/satryanararya/go-chefbot/utils/error"
	rand_util "github.com/satryanararya/go-chefbot/utils/recommendation"
)

type RecipeClient interface {
	SearchRecipe(ctx context.Context, name string) (SearchRecipeResponse, error)
	GetRecipeInformation(ctx context.Context, recipeID int) (RecipeInformation, error)
	GetMultipleRecipeInformation(ctx context.Context, recipeIDs []int) ([]RecipeInformation, error)
}

type recipeClient struct {
	APIKey string
	Client *http.Client
}

func NewRecipeClient(apiKey string) *recipeClient {
	return &recipeClient{
		APIKey: apiKey,
		Client: http.DefaultClient,
	}
}

func (r *recipeClient) SearchRecipe(ctx context.Context, name string) (SearchRecipeResponse, error) {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?query=%s&number=%d", name, rand_util.GenerateRandomNumber())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return SearchRecipeResponse{}, err_util.ErrExternalService
	}

	req.Header.Set("X-API-Key", r.APIKey)

	fmt.Println("name", name)

	resp, err := r.Client.Do(req)
	if err != nil {
		fmt.Println("err", err)
		return SearchRecipeResponse{}, err_util.ErrExternalService
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return SearchRecipeResponse{}, err_util.ErrItemNotFound
	}

	searchRecipeRes := new(SearchRecipeResponse)
	err = json.NewDecoder(resp.Body).Decode(searchRecipeRes)
	fmt.Println("searchRecipeRes", searchRecipeRes)
	if err != nil {
		return SearchRecipeResponse{}, err
	}

	return *searchRecipeRes, nil
}

func (r *recipeClient) GetRecipeInformation(ctx context.Context, recipeID int) (RecipeInformation, error) {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/%d/information", recipeID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return RecipeInformation{}, err_util.ErrExternalService
	}

	req.Header.Set("X-API-Key", r.APIKey)

	resp, err := r.Client.Do(req)
	if err != nil {
		return RecipeInformation{}, err_util.ErrExternalService
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return RecipeInformation{}, err_util.ErrItemNotFound
	}

	var recipeInformationResponse RecipeInformation
	err = json.NewDecoder(resp.Body).Decode(&recipeInformationResponse)
	if err != nil {
		return RecipeInformation{}, err
	}

	return recipeInformationResponse, nil
}

func (r *recipeClient) GetMultipleRecipeInformation(ctx context.Context, recipeIDs []int) ([]RecipeInformation, error) {
	var recipeInformations []RecipeInformation

	for _, recipeID := range recipeIDs {
		recipeInformation, err := r.GetRecipeInformation(ctx, recipeID)
		if err != nil {
			return nil, err
		}

		recipeInformations = append(recipeInformations, recipeInformation)
	}

	return recipeInformations, nil
}