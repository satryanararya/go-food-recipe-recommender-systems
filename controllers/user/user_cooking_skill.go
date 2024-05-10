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

type userCookingSkillController struct {
	userCookingSkillUsecase usecases.UserCookingSkillUseCase
	validator               *validation.Validator
	token                   token.TokenUtil
}

func NewUserCookingSkillController(ucsc usecases.UserCookingSkillUseCase, v *validation.Validator, t token.TokenUtil) *userCookingSkillController {
	return &userCookingSkillController{
		userCookingSkillUsecase: ucsc,
		validator:               v,
		token:                   t,
	}
}

func (ucsc *userCookingSkillController) AddUserCookingSkill(c echo.Context) error {
	claims := ucsc.token.GetClaims(c)
	var req = new(dto.UserCookingSkillRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := ucsc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	err := ucsc.userCookingSkillUsecase.AddCookingSkill(c, claims.ID, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgAddCookingSkillFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgAddCookingSkillSuccess, nil)
}

func (ucsc *userCookingSkillController) EditUserCookingSkill(c echo.Context) error {
	claims := ucsc.token.GetClaims(c)
	var req = new(dto.UserCookingSkillRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := ucsc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}
	err := ucsc.userCookingSkillUsecase.EditCookingSkill(c, claims.ID, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgEditCookingSkillFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgEditCookingSkillSuccess, nil)
}

