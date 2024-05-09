package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	msg "github.com/satryanararya/go-chefbot/constants/message"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/usecases"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
)

type userFoodPreferencesController struct {
	userFoodPreferencesUseCase usecases.UserFoodPreferencesUseCase
	validator                  *validation.Validator
	token                      token.TokenUtil
}

func NewUserFoodPreferencesController(ufpc usecases.UserFoodPreferencesUseCase, v *validation.Validator, t token.TokenUtil) *userFoodPreferencesController {
	return &userFoodPreferencesController{
		userFoodPreferencesUseCase: ufpc,
		validator:                  v,
		token:                      t,
	}
}

func (ufpc *userFoodPreferencesController) AddUserFoodPreferences(c echo.Context) error {
	claims := ufpc.token.GetClaims(c)
	var req = new(dto.UserFoodPreferencesRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := ufpc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}

	err := ufpc.userFoodPreferencesUseCase.AddFoodPreference(c, claims.ID, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgAddPreferenceFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgAddPreferenceSuccess, nil)
}

func (ufpc *userFoodPreferencesController) EditUserFoodPreferences(c echo.Context) error {
    claims := ufpc.token.GetClaims(c)
    var req = new(dto.UserFoodPreferencesRequest)
    if err := c.Bind(req); err != nil {
        return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
    }
    if err := ufpc.validator.Validate(req); err != nil {
        return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
    }

    err := ufpc.userFoodPreferencesUseCase.EditFoodPreference(c, claims.ID, req)
    if err != nil {
        return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgEditPreferenceFailed)
    }
    return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgEditPreferenceSuccess, nil)
}

func (ufpc *userFoodPreferencesController) DeleteUserFoodPreferences(c echo.Context) error {
    claims := ufpc.token.GetClaims(c)

    err := ufpc.userFoodPreferencesUseCase.DeleteFoodPreference(c, claims.ID)
    if err != nil {
        return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgDeletePreferenceFailed)
    }
    return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgDeletePreferenceSuccess, nil)
}