package http

import (
	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/base"
	status "github.com/satryanararya/go-chefbot/constants/status"
)

func HandleErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, &dto.BaseResponse{
		Status:  status.StatusFailed,
		Message: message,
	})
}

func HandleSuccessResponse(c echo.Context, code int, message string, data any) error {
	return c.JSON(code, &dto.BaseResponse{
		Status:  status.StatusSuccess,
		Message: message,
		Data:    data,
	})
}