package recipe

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	msg "github.com/satryanararya/go-chefbot/constants/message"
	dto "github.com/satryanararya/go-chefbot/dto/recipe"
	"github.com/satryanararya/go-chefbot/usecases"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
)

type recipeController struct {
	recipeUseCase usecases.RecipeUseCase
	validator     *validation.Validator
	token         token.TokenUtil
}

func NewRecipeController(rc usecases.RecipeUseCase, v *validation.Validator, t token.TokenUtil) *recipeController {
	return &recipeController{
		recipeUseCase: rc,
		validator:     v,
		token:         t,
	}
}

func (rc *recipeController) CreateRecipe(c echo.Context) error {
	claims := rc.token.GetClaims(c)
	var req = new(dto.CreateRecipeRequest)

	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := rc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	res, err := rc.recipeUseCase.CreateRecipe(c, claims.ID, req)
	if err != nil {
		fmt.Println(err)
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgAddPreferenceFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgAddPreferenceSuccess, res)
}

func (rc *recipeController) UpdateRecipe(c echo.Context) error {
	claims := rc.token.GetClaims(c)
	var req = new(dto.UpdateRecipeRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := rc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	err := rc.recipeUseCase.UpdateRecipe(c, claims.ID, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgEditPreferenceFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgEditPreferenceSuccess, nil)
}

func (rc *recipeController) GetRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidID)
	}
	res, err := rc.recipeUseCase.GetRecipe(c, int64(id))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgGetRecipeFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgGetRecipeSuccess, res)
}

func (rc *recipeController) DeleteRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidID)
	}
	err = rc.recipeUseCase.DeleteRecipe(c, int64(id))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgDeleteRecipeFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgDeleteRecipeSuccess, nil)
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

