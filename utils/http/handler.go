package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	status "github.com/satryanararya/go-chefbot/constants/status"
	"github.com/satryanararya/go-chefbot/dto"
	dto_base "github.com/satryanararya/go-chefbot/dto/base"
)

func HandleErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, &dto_base.BaseResponse{
		Status:  status.StatusFailed,
		Message: message,
	})
}

func HandleSuccessResponse(c echo.Context, code int, message string, data any) error {
	return c.JSON(code, &dto_base.BaseResponse{
		Status:  status.StatusSuccess,
		Message: message,
		Data:    data,
	})
}

func HandlePaginationResponse(
	c echo.Context,
	message string,
	data any,
	pagination *dto.PaginationMetadata,
	link *dto.Link,
) error {
	return c.JSON(http.StatusOK, &dto.PaginationResponse{
		BaseResponse: dto_base.BaseResponse{
			Status:  status.StatusSuccess,
			Message: message,
			Data:    data,
		},
		Pagination: pagination,
		Link:       link,
	})
}
