package usecases_test

import (
	"context"
	"time"
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_repo "github.com/satryanararya/go-chefbot/mocks/repositories"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/stretchr/testify/assert"
)

func TestNewFavoriteRecipeUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewFavoriteRecipeUsecase(mock_repo.NewMockFavoriteRecipeRepository(t)),
	)
}

func TestAddToFavorites(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/favorites_recipe/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockFavoriteRecipeRepo := new(mock_repo.MockFavoriteRecipeRepository)
	mockFavoriteRecipeRepo.On("AddToFavorites", ctx, &entities.FavoriteRecipe{UserID: 1, RecipeID: 123}).Return(nil)

	favoriteRecipeUsecase := usecases.NewFavoriteRecipeUsecase(mockFavoriteRecipeRepo)
	err := favoriteRecipeUsecase.AddToFavorites(c, 1, 123)
	assert.NoError(t, err)
}

func TestRemoveFromFavorites(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/favorites_recipe/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockFavoriteRecipeRepo := new(mock_repo.MockFavoriteRecipeRepository)
	mockFavoriteRecipeRepo.On("RemoveFromFavorites", ctx, int64(1), int64(123)).Return(nil)

	favoriteRecipeUsecase := usecases.NewFavoriteRecipeUsecase(mockFavoriteRecipeRepo)
	err := favoriteRecipeUsecase.RemoveFromFavorites(c, 1, 123)
	assert.NoError(t, err)
}

func TestFindFavoritesByUserID(t *testing.T) {
	example := &[]entities.FavoriteRecipe{
		{
			ID:          1,
			UserID:      1,
			RecipeID:    123,
			RecipeTitle: "Recipe 1",
			CreatedAt:   time.UnixMilli(1714757476909),
			UpdatedAt:   time.UnixMilli(1714757476909),
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/favorites", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockFavoriteRecipeRepo := new(mock_repo.MockFavoriteRecipeRepository)
	mockFavoriteRecipeRepo.On("FindFavoritesByUserID", ctx, int64(1)).Return(*example, nil)
	favoriteRecipeUsecase := usecases.NewFavoriteRecipeUsecase(mockFavoriteRecipeRepo)
	res, err := favoriteRecipeUsecase.FindFavoritesByUserID(c, 1)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
