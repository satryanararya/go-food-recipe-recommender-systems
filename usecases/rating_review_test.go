package usecases_test

import (
	"context"
	"time"

	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	dto "github.com/satryanararya/go-chefbot/dto/rating_review"
	mock_repo "github.com/satryanararya/go-chefbot/mocks/repositories"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/stretchr/testify/assert"
)

func TestNewRatingReviewUseCase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewRatingReviewUseCase(mock_repo.NewMockRatingReviewRepository(t)),
	)
}

func TestCreateRatingReview(t *testing.T) {
	r := &dto.RatingReviewRequest{
		Rating: 5,
		Review: "Great",
	}
	d := &entities.RatingReview{
		UserID: 1,
		RecipeID: 123,
		Rating: 5,
		Review: "Great",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/rating_review/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockRatingReviewRepo := new(mock_repo.MockRatingReviewRepository)
	mockRatingReviewRepo.On("Create", ctx, d).Return(nil)

	ratingReviewUsecase := usecases.NewRatingReviewUseCase(mockRatingReviewRepo)
	err := ratingReviewUsecase.CreateRatingReview(c, 1, 123, r)
	assert.NoError(t, err)
}

func TestDeleteRatingReview(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/rating_review/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockRatingReviewRepo := new(mock_repo.MockRatingReviewRepository)
	mockRatingReviewRepo.On("Delete", ctx, int64(1), int64(123)).Return(nil)

	ratingReviewUsecase := usecases.NewRatingReviewUseCase(mockRatingReviewRepo)
	err := ratingReviewUsecase.DeleteRatingReview(c, 1, 123)
	assert.NoError(t, err)
}

func TestGetUserRatingReviews(t *testing.T) {
	example := &[]entities.RatingReview{
		{
			ID:          1,
			UserID:      1,
			RecipeID:    123,
			Rating:      5,
			Review:      "Great",
			CreatedAt:   time.UnixMilli(1714757476909),
			UpdatedAt:   time.UnixMilli(1714757476909),
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reviews", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockRatingReviewRepo := new(mock_repo.MockRatingReviewRepository)
	mockRatingReviewRepo.On("FindByUserID", ctx, int64(1)).Return(*example, nil)

	ratingReviewUsecase := usecases.NewRatingReviewUseCase(mockRatingReviewRepo)
	_, err := ratingReviewUsecase.GetUserRatingReviews(c, 1)
	assert.NoError(t, err)
}