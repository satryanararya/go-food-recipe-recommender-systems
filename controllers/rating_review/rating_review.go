package ratingreview

import (
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/rating_review"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	msg "github.com/satryanararya/go-chefbot/constants/message"
)

type ratingReviewController struct {
	ratingReviewUseCase usecases.RatingReviewUseCase
	validator           *validation.Validator
	token               token.TokenUtil
}

func NewRatingReviewController(rruc usecases.RatingReviewUseCase, v *validation.Validator, t token.TokenUtil) *ratingReviewController {
	return &ratingReviewController{
		ratingReviewUseCase: rruc,
		validator:           v,
		token:               t,
	}
}

func (rruc *ratingReviewController) CreateRatingReview(c echo.Context) error {
	claims := rruc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var req = new(dto.RatingReviewRequest)
	if err := c.Bind(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgMismatchedDataType)
	}
	if err := rruc.validator.Validate(req); err != nil {
		return http_util.HandleErrorResponse(c, http.StatusBadRequest, msg.MsgInvalidRequestData)
	}

	err = rruc.ratingReviewUseCase.CreateRatingReview(c, claims.ID, int64(recipeID), req)
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgAddRatingReviewFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusCreated, msg.MsgAddRatingReviewSuccess, nil)
}

func (rruc *ratingReviewController) DeleteRatingReview(c echo.Context) error {
	claims := rruc.token.GetClaims(c)
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = rruc.ratingReviewUseCase.DeleteRatingReview(c, claims.ID, int64(recipeID))
	if err != nil {
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgDeleteRatingReviewFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgDeleteRatingReviewSuccess, nil)
}

func (rruc *ratingReviewController) GetUserRatingReviews(c echo.Context) error {
    claims := rruc.token.GetClaims(c)

    ratingReviews, err := rruc.ratingReviewUseCase.GetUserRatingReviews(c, claims.ID)
    if err != nil {
        return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgGetRatingReviewFailed)
    }

    return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgGetRatingReviewSuccess, ratingReviews)
}