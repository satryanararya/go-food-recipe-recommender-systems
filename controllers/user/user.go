package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/usecases"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	"github.com/satryanararya/go-chefbot/utils/validation"
	msg "github.com/satryanararya/go-chefbot/constants/message"
)

type userController struct {
	userUseCase usecases.UserUseCase
	validator   *validation.Validator
}

func NewUserController(uuc usecases.UserUseCase, v *validation.Validator) *userController {
	return &userController{
		userUseCase: uuc,
		validator:   v,
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

	if err := uc.userUseCase.Register(c, req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgUserCreationFailed)
	}

	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgUserCreated, nil)
}
