package ratingreview

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	rc "github.com/satryanararya/go-chefbot/controllers/rating_review"
	rr "github.com/satryanararya/go-chefbot/repositories"
	ru "github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

func InitRatingReviewRoute(rrg *echo.Group, db *gorm.DB, v *validation.Validator) {
	ratingReviewRepo := rr.NewRatingReviewRepository(db)
	ratingReviewUseCase := ru.NewRatingReviewUseCase(ratingReviewRepo)

	tokenUtil := token.NewTokenUtil()

	ratingReviewController := rc.NewRatingReviewController(ratingReviewUseCase, v, tokenUtil)

    rrg.Use(echojwt.WithConfig(token.GetJWTConfig()))

	rrg.POST("/rating_review/:id", ratingReviewController.CreateRatingReview)
	rrg.DELETE("/rating_review/:id", ratingReviewController.DeleteRatingReview)
	rrg.GET("/reviews", ratingReviewController.GetUserRatingReviews)
}