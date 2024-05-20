package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/constants/enums"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
	"github.com/satryanararya/go-chefbot/utils/validation"
)

type UserCookingSkillUseCase interface {
	AddCookingSkill(c echo.Context, userID uuid.UUID, req *dto.UserCookingSkillRequest) error
	EditCookingSkill(c echo.Context, userID uuid.UUID, req *dto.UserCookingSkillRequest) error
}

type userCookingSkillUseCase struct {
	userCookingSkillRepo repositories.UserCookingSkillRepository
}

func NewUserCookingSkillUseCase(repo repositories.UserCookingSkillRepository) *userCookingSkillUseCase {
	return &userCookingSkillUseCase{
		userCookingSkillRepo: repo,
	}
}

func (uc *userCookingSkillUseCase) AddCookingSkill(c echo.Context, userID uuid.UUID, req *dto.UserCookingSkillRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	if err := uc.validateUserCookingSkillRequest(req); err != nil {
		return err
	}

	userCookingSkill := &entities.UserCookingSkill{
		UserID:              userID,
		ExperienceYears:     string(enums.ExperienceYears(req.ExperienceYears)),
		TimeCommitment:      string(enums.TimeCommitment(req.TimeCommitment)),
		RecipeComplexity:    string(enums.RecipeComplexity(req.RecipeComplexity)),
		IngredientDiversity: string(enums.IngredientDiversity(req.IngredientDiversity)),
	}

	return uc.userCookingSkillRepo.AddCookingSkill(ctx, userCookingSkill)
}

func (uc *userCookingSkillUseCase) EditCookingSkill(c echo.Context, userID uuid.UUID, req *dto.UserCookingSkillRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	if err := uc.validateUserCookingSkillRequest(req); err != nil {
		return err
	}

	userCookingSkill := &entities.UserCookingSkill{
		UserID:              userID,
		ExperienceYears:     string(enums.ExperienceYears(req.ExperienceYears)),
		TimeCommitment:      string(enums.TimeCommitment(req.TimeCommitment)),
		RecipeComplexity:    string(enums.RecipeComplexity(req.RecipeComplexity)),
		IngredientDiversity: string(enums.IngredientDiversity(req.IngredientDiversity)),
	}

	return uc.userCookingSkillRepo.EditCookingSkill(ctx, userCookingSkill)
}

func (uc *userCookingSkillUseCase) validateUserCookingSkillRequest(req *dto.UserCookingSkillRequest) error {
    if req.ExperienceYears != "" && !validation.IsValidEnumValue("ExperienceYears", req.ExperienceYears) {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ExperienceYears")
    }
    if req.TimeCommitment != "" && !validation.IsValidEnumValue("TimeCommitment", req.TimeCommitment) {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid TimeCommitment")
    }
    if req.RecipeComplexity != "" && !validation.IsValidEnumValue("RecipeComplexity", req.RecipeComplexity) {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid RecipeComplexity")
    }
    if req.IngredientDiversity != "" && !validation.IsValidEnumValue("IngredientDiversity", req.IngredientDiversity) {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid IngredientDiversity")
    }
    return nil
}