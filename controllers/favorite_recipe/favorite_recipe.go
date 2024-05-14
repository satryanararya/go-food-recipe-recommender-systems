package favoriterecipe

import (
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	msg "github.com/satryanararya/go-chefbot/constants/message"
)

type favoriteRecipeController struct {
	favoriteRecipeUseCase usecases.FavoriteRecipeUsecase
	validator             *validation.Validator
	token                 token.TokenUtil
}

func NewFavoriteRecipeController(fruc usecases.FavoriteRecipeUsecase, v *validation.Validator, t token.TokenUtil) *favoriteRecipeController {
	return &favoriteRecipeController{
		favoriteRecipeUseCase: fruc,
		validator:             v,
		token:                 t,
	}
}

func (fruc *favoriteRecipeController) AddToFavorites(c echo.Context) error {
	claims := fruc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = fruc.favoriteRecipeUseCase.AddToFavorites(c, claims.ID, int64(recipeID))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgAddFavoriteRecipeFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgAddFavoriteRecipeSuccess, nil)
}

func (fruc *favoriteRecipeController) RemoveFromFavorites(c echo.Context) error {
	claims := fruc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = fruc.favoriteRecipeUseCase.RemoveFromFavorites(c, claims.ID, int64(recipeID))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgRemoveFavoriteRecipeFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgRemoveFavoriteRecipeSuccess, nil)
}

func (fruc *favoriteRecipeController) FindFavoritesByUserID(c echo.Context) error {
	claims := fruc.token.GetClaims(c)

	favoriteRecipes, err := fruc.favoriteRecipeUseCase.FindFavoritesByUserID(c, claims.ID)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgGetFavoriteRecipesFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgGetFavoriteRecipesSuccess, favoriteRecipes)
}