package usecases

import (
	"context"

	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/rating_review"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
)

type RatingReviewUseCase interface {
	CreateRatingReview(c echo.Context, userID int64, recipeID int64, req *dto.RatingReviewRequest) error
	DeleteRatingReview(c echo.Context, userID int64, recipeID int64) error
	GetUserRatingReviews(c echo.Context, userID int64) ([]dto.RatingReviewResponse, error)
}

type ratingReviewUseCase struct {
	ratingReviewRepo repositories.RatingReviewRepository
}

func NewRatingReviewUseCase(rr repositories.RatingReviewRepository) *ratingReviewUseCase {
	return &ratingReviewUseCase{
		ratingReviewRepo: rr,
	}
}

func (rruc *ratingReviewUseCase) CreateRatingReview(c echo.Context, userID int64, recipeID int64, req *dto.RatingReviewRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	ratingReview := &entities.RatingReview{
		UserID:   userID,
		RecipeID: recipeID,
		Rating:   req.Rating,
		Review:   req.Review,
	}
	return rruc.ratingReviewRepo.Create(ctx, ratingReview)
}

func (rruc *ratingReviewUseCase) DeleteRatingReview(c echo.Context, userID int64, recipeID int64) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return rruc.ratingReviewRepo.Delete(ctx, userID, recipeID)
}

func (rruc *ratingReviewUseCase) GetUserRatingReviews(c echo.Context, userID int64) ([]dto.RatingReviewResponse, error) {
    ctx, cancel := context.WithCancel(c.Request().Context())
    defer cancel()

    ratingReviews, err := rruc.ratingReviewRepo.FindByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }

    var responses []dto.RatingReviewResponse
    for _, rr := range ratingReviews {
        responses = append(responses, dto.RatingReviewResponse{
            UserID:   rr.UserID,
            RecipeID: rr.RecipeID,
            Rating:   rr.Rating,
            Review:   rr.Review,
        })
    }

    return responses, nil
}