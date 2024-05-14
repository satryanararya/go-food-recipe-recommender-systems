package repositories

import (
	"context"

	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type RatingReviewRepository interface {
	Create(ctx context.Context, ratingReview *entities.RatingReview) error
	Delete(ctx context.Context, userID int64, recipeID int64) error
	FindByUserID(ctx context.Context, userID int64) ([]entities.RatingReview, error)
}

type ratingReviewRepository struct {
	DB *gorm.DB
}

func NewRatingReviewRepository(db *gorm.DB) *ratingReviewRepository {
	return &ratingReviewRepository{
		DB: db,
	}
}

func (r *ratingReviewRepository) Create(ctx context.Context, ratingReview *entities.RatingReview) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.DB.Create(ratingReview).Error
}

func (r *ratingReviewRepository) Delete(ctx context.Context, userID int64, recipeID int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.DB.Where("user_id = ? AND recipe_id = ?", userID, recipeID).Delete(&entities.RatingReview{}).Error
}

func (r *ratingReviewRepository) FindByUserID(ctx context.Context, userID int64) ([]entities.RatingReview, error) {
    if err := ctx.Err(); err != nil {
        return nil, err
    }

    var ratingReviews []entities.RatingReview
    if err := r.DB.Where("user_id = ?", userID).Find(&ratingReviews).Error; err != nil {
        return nil, err
    }

    return ratingReviews, nil
}