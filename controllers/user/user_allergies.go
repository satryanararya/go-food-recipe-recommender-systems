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

type userAllergiesController struct {
	userAllergiesUseCase usecases.UserAllergiesUseCase
	validator            *validation.Validator
	token                token.TokenUtil
}

func NewUserAllergiesController(uac usecases.UserAllergiesUseCase, v *validation.Validator, t token.TokenUtil) *userAllergiesController {
	return &userAllergiesController{
		userAllergiesUseCase: uac,
		validator:            v,
		token:                t,
	}
}

func (uac *userAllergiesController) AddUserAllergy(c echo.Context) error {
	claims := uac.token.GetClaims(c)
	var req = new(dto.UserAllergiesRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := uac.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	res, err := uac.userAllergiesUseCase.GetIngredientInfo(c.Request().Context(), claims.ID, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgAddUserAllergyFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgAddUserAllergySuccess, res)
}
