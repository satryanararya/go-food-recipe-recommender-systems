package usecases

import (
	"context"

	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type UserFoodPreferencesUseCase interface {
	AddFoodPreference(c echo.Context, id int64, req *dto.UserFoodPreferencesRequest) error
	EditFoodPreference(c echo.Context, id int64, req *dto.UserFoodPreferencesRequest) error
    DeleteFoodPreference(c echo.Context, id int64) error
}

type userFoodPreferencesUseCase struct {
	userFoodPrefRepo repositories.UserFoodPreferencesRepository
}

func NewUserFoodPreferencesUseCase(ufpr repositories.UserFoodPreferencesRepository) *userFoodPreferencesUseCase {
	return &userFoodPreferencesUseCase{
		userFoodPrefRepo: ufpr,
	}
}

func (upc *userFoodPreferencesUseCase) AddFoodPreference(c echo.Context, id int64, req *dto.UserFoodPreferencesRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	userFoodPref := &entities.UserFoodPreference{
		UserID: id,
		DietaryRestriction: req.DietaryRestriction,
		ReligiousReason: req.ReligiousReason,
	}
	return upc.userFoodPrefRepo.AddFoodPreference(ctx, userFoodPref)
}

func (upc *userFoodPreferencesUseCase) EditFoodPreference(c echo.Context, id int64, req *dto.UserFoodPreferencesRequest) error {
    ctx, cancel := context.WithCancel(c.Request().Context())
    defer cancel()

    userFoodPref := &entities.UserFoodPreference{
        UserID:             id,
        DietaryRestriction: req.DietaryRestriction,
        ReligiousReason:    req.ReligiousReason,
    }
    return upc.userFoodPrefRepo.EditFoodPreference(ctx, userFoodPref)
}

func (upc *userFoodPreferencesUseCase) DeleteFoodPreference(c echo.Context, id int64) error {
    ctx, cancel := context.WithCancel(c.Request().Context())
    defer cancel()

    return upc.userFoodPrefRepo.DeleteFoodPreference(ctx, id)
}