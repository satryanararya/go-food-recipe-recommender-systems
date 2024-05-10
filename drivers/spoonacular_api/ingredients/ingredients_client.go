package ingredients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/satryanararya/go-chefbot/dto/user"
)

type IngredientClient interface {
	SearchIngredient(ctx context.Context, dto *user.UserAllergiesRequest) (int, error)
	GetIngredient(ctx context.Context, ingredientID int) (*Ingredient, error)
}

type ingredientClient struct {
	APIKey  string
	Client  *http.Client
}

func NewIngredientClient(apiKey string) *ingredientClient {
	return &ingredientClient{
		APIKey:  apiKey,
		Client:  http.DefaultClient,
	}
}

func (c *ingredientClient) SearchIngredient(ctx context.Context, dto *user.UserAllergiesRequest) (int, error) {
	url := fmt.Sprintf("https://api.spoonacular.com/food/ingredients/search?query=%s&number=1", dto.IngredientName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		if err == context.Canceled {
			return 0, err
		}
		return 0, errors.Wrap(err, "failed to do request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var searchResponse SearchIngredientResponse
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return 0, errors.Wrap(err, "failed to decode response")
	}

	if len(searchResponse.Results) > 0 {
		return searchResponse.Results[0].ID, nil
	}

	return 0, fmt.Errorf("no results found for ingredient: %s", dto.IngredientName)
}

func (c *ingredientClient) GetIngredient(ctx context.Context, ingredientID int) (*Ingredient, error) {
	url := fmt.Sprintf("https://api.spoonacular.com/food/ingredients/%d/information", ingredientID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		if err == context.Canceled {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed to do request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var ingredient Ingredient
	err = json.NewDecoder(resp.Body).Decode(&ingredient)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return &ingredient, nil
}
