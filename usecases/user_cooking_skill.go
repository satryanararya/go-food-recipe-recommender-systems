package usecases

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/constants/enums"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
	"github.com/satryanararya/go-chefbot/utils/validation"
)

type UserCookingSkillUseCase interface {
	AddCookingSkill(c echo.Context, userID int64, req *dto.UserCookingSkillRequest) error
	EditCookingSkill(c echo.Context, userID int64, req *dto.UserCookingSkillRequest) error
}

type userCookingSkillUseCase struct {
	userCookingSkillRepo repositories.UserCookingSkillRepository
}

func NewUserCookingSkillUseCase(repo repositories.UserCookingSkillRepository) *userCookingSkillUseCase {
	return &userCookingSkillUseCase{
		userCookingSkillRepo: repo,
	}
}

func (uc *userCookingSkillUseCase) AddCookingSkill(c echo.Context, userID int64, req *dto.UserCookingSkillRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	if err := uc.ValidateUserCookingSkillRequest(req); err != nil {
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

func (uc *userCookingSkillUseCase) EditCookingSkill(c echo.Context, userID int64, req *dto.UserCookingSkillRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	if err := uc.ValidateUserCookingSkillRequest(req); err != nil {
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

func (uc *userCookingSkillUseCase) ValidateUserCookingSkillRequest(req *dto.UserCookingSkillRequest) error {
	if !validation.IsValidEnumValue("ExperienceYears", req.ExperienceYears) ||
		!validation.IsValidEnumValue("TimeCommitment", req.TimeCommitment) ||
		!validation.IsValidEnumValue("RecipeComplexity", req.RecipeComplexity) ||
		!validation.IsValidEnumValue("IngredientDiversity", req.IngredientDiversity) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	return nil
}
