package recipe

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	// "github.com/google/uuid"
	"github.com/labstack/echo/v4"

	msg "github.com/satryanararya/go-chefbot/constants/message"
	"github.com/satryanararya/go-chefbot/dto"
	dto_recipe "github.com/satryanararya/go-chefbot/dto/recipe"
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
	var req = new(dto_recipe.RecipeRequest)

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

func (rc *recipeController) UploadRecipeImage(c echo.Context) error {
	claims := rc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("recipeID"))
	var req = new(dto_recipe.RecipeImageRequest)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, "Invalid recipe ID")
	}

	err = rc.recipeUseCase.UploadRecipeImage(c, claims.ID, recipeID, req)
	if err != nil {
		fmt.Println(err)
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, "Failed to upload image")
	}

	return http_util.HandleSuccessResponse(c, http.StatusOK, "Image uploaded successfully", nil)
}

func (rc *recipeController) GetUserRecipes(c echo.Context) error {
	claims := rc.token.GetClaims(c)

	page := strings.TrimSpace(c.QueryParam("page"))
	limit := strings.TrimSpace(c.QueryParam("limit"))

	intPage, intLimit, err := rc.convertQueryParams(page, limit)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	p := &dto.PaginationRequest{
		Page:  intPage,
		Limit: intLimit,
	}
	if err := rc.validator.Validate(p); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	res, meta, link, err := rc.recipeUseCase.GetUserRecipes(c, claims.ID, p)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgGetRecipeFailed)
	}
	return http_util.HandlePaginationResponse(c, msg.MsgGetRecipeSuccess, res, meta, link)
}

func (rc *recipeController) UpdateRecipe(c echo.Context) error {
	claims := rc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("recipeID"))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidID)
	}
	var req = new(dto_recipe.RecipeRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := rc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	res, err := rc.recipeUseCase.UpdateRecipe(c, claims.ID, recipeID, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgUpdateRecipeFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgUpdateRecipeSuccess, res)
}

func (rc *recipeController) DeleteRecipe(c echo.Context) error {
	claims := rc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("recipeID"))

	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidID)
	}

	err = rc.recipeUseCase.DeleteRecipe(c, claims.ID, recipeID)
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

func (rc *recipeController) convertQueryParams(page, limit string) (int, int, error) {
	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	var (
		intPage, intLimit int
		err               error
	)

	intPage, err = strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	intLimit, err = strconv.Atoi(limit)
	if err != nil {
		return 0, 0, err
	}

	return intPage, intLimit, nil
}
