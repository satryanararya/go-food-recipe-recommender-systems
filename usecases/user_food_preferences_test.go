package usecases_test

import (
	"context"

	// "time"
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	// "gorm.io/gorm"

	mock_repo "github.com/satryanararya/go-chefbot/mocks/repositories"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/stretchr/testify/assert"
)

func TestNewUserFoodPreferencesUseCase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewUserFoodPreferencesUseCase(mock_repo.NewMockUserFoodPreferencesRepository(t)),
	)
}

func TestAddFoodPreference(t *testing.T) {
	dietaryRestriction := "Vegetarian"
	religiousReason := "Halal"

	r := &dto.UserFoodPreferencesRequest{
		DietaryRestriction: &dietaryRestriction,
		ReligiousReason:    &religiousReason,
	}
	uid := uuid.New()
	d := &entities.UserFoodPreference{
		UserID:             uid,
		DietaryRestriction: &dietaryRestriction,
		ReligiousReason:    &religiousReason,
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/food_preferences/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockUserFoodPreferencesRepo := new(mock_repo.MockUserFoodPreferencesRepository)
	mockUserFoodPreferencesRepo.On("AddFoodPreference", ctx, d).Return(nil)

	userFoodPreferencesUsecase := usecases.NewUserFoodPreferencesUseCase(mockUserFoodPreferencesRepo)
	err := userFoodPreferencesUsecase.AddFoodPreference(c, uid, r)
	assert.NoError(t, err)
}

func TestEditFoodPreference(t *testing.T) {
	dietaryRestriction := "Vegetarian"
	religiousReason := "Halal"

	r := &dto.UserFoodPreferencesRequest{
		DietaryRestriction: &dietaryRestriction,
		ReligiousReason:    &religiousReason,
	}
	uid := uuid.New()
	d := &entities.UserFoodPreference{
		UserID:             uid,
		DietaryRestriction: &dietaryRestriction,
		ReligiousReason:    &religiousReason,
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/food_preferences/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockUserFoodPreferencesRepo := new(mock_repo.MockUserFoodPreferencesRepository)
	mockUserFoodPreferencesRepo.On("EditFoodPreference", ctx, d).Return(nil)

	userFoodPreferencesUsecase := usecases.NewUserFoodPreferencesUseCase(mockUserFoodPreferencesRepo)
	err := userFoodPreferencesUsecase.EditFoodPreference(c, uid, r)
	assert.NoError(t, err)
}

func TestDeleteFoodPreference(t *testing.T) {
	uid := uuid.New()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/food_preferences/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockUserFoodPreferencesRepo := new(mock_repo.MockUserFoodPreferencesRepository)
	mockUserFoodPreferencesRepo.On("DeleteFoodPreference", ctx, uid).Return(nil)

	userFoodPreferencesUsecase := usecases.NewUserFoodPreferencesUseCase(mockUserFoodPreferencesRepo)
	err := userFoodPreferencesUsecase.DeleteFoodPreference(c, uid)
	assert.NoError(t, err)
}
