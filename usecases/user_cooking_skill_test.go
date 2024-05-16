package usecases_test

import (
	"context"
	// "time"
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	dto "github.com/satryanararya/go-chefbot/dto/user"

	mock_repo "github.com/satryanararya/go-chefbot/mocks/repositories"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/stretchr/testify/assert"
)

func TestNewUserCookingSkill(t *testing.T){
	assert.NotNil(
		t,
		usecases.NewUserCookingSkillUseCase(mock_repo.NewMockUserCookingSkillRepository(t)),
	)
}

func TestAddCookingSkill(t *testing.T){
	r := &dto.UserCookingSkillRequest{
		ExperienceYears: "0-2 tahun",
		TimeCommitment: "1-2x per minggu",
		RecipeComplexity: "sederhana",
		IngredientDiversity: "jarang",
	}
	d := &entities.UserCookingSkill{
		UserID: 1,
		ExperienceYears: "0-2 tahun",
		TimeCommitment: "1-2x per minggu",
		RecipeComplexity: "sederhana",
		IngredientDiversity: "jarang",
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/cooking_skill/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockUserCookingSkillRepo := new(mock_repo.MockUserCookingSkillRepository)
	mockUserCookingSkillRepo.On("AddCookingSkill", ctx, d).Return(nil)

	userCookingSkillUsecase := usecases.NewUserCookingSkillUseCase(mockUserCookingSkillRepo)
	err := userCookingSkillUsecase.AddCookingSkill(c, 1, r)
	assert.NoError(t, err)
}

func TestEditCookingSkill(t *testing.T){
	r := &dto.UserCookingSkillRequest{
		ExperienceYears: "0-2 tahun",
		TimeCommitment: "1-2x per minggu",
		RecipeComplexity: "sederhana",
		IngredientDiversity: "jarang",
	}
	d := &entities.UserCookingSkill{
		UserID: 1,
		ExperienceYears: "0-2 tahun",
		TimeCommitment: "1-2x per minggu",
		RecipeComplexity: "sederhana",
		IngredientDiversity: "jarang",
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/cooking_skill/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockUserCookingSkillRepo := new(mock_repo.MockUserCookingSkillRepository)
	mockUserCookingSkillRepo.On("EditCookingSkill", ctx, d).Return(nil)

	userCookingSkillUsecase := usecases.NewUserCookingSkillUseCase(mockUserCookingSkillRepo)
	err := userCookingSkillUsecase.EditCookingSkill(c, 1, r)
	assert.NoError(t, err)
}