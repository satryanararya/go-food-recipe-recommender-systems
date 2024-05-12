package recipe

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/validation"
)

type recipeController struct {
	recipeUseCase   usecases.RecipeUseCase
	validator *validation.Validator
}

func NewRecipeController(useCase usecases.RecipeUseCase, v *validation.Validator) *recipeController {
	return &recipeController{
		recipeUseCase:   useCase,
		validator: v,
	}
}

func (rc *recipeController) SearchRecipe(c echo.Context) error {
	name := c.QueryParam("name")
	recipes, err := rc.recipeUseCase.SearchRecipe(c, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, recipes)
}

func (rc *recipeController) GetRecipeInformation(c echo.Context) error {
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid recipe ID")
	}
	recipeInformation, err := rc.recipeUseCase.GetRecipeInformation(c, recipeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, recipeInformation)
}
