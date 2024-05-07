package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	msg "github.com/satryanararya/go-chefbot/constants/message"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/usecases"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"github.com/satryanararya/go-chefbot/utils/token"
)

type userController struct {
	userUseCase usecases.UserUseCase
	validator   *validation.Validator
	token       token.TokenUtil
}

func NewUserController(uuc usecases.UserUseCase, v *validation.Validator, t token.TokenUtil) *userController {
	return &userController{
		userUseCase: uuc,
		validator:   v,
		token: t,
	}
}

func (uc *userController) Register(c echo.Context) error {
	var req = new(dto.UserRegisterRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}

	if err := uc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}

	res, err := uc.userUseCase.Register(c, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgUserCreationFailed)
	}

	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgUserCreated, res)
}

func (uc *userController) Login(c echo.Context) error {
	var req = new(dto.UserLoginRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}

	if err := uc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}

	res, err := uc.userUseCase.Login(c, req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgLoginFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgLoginSuccess, res)
}