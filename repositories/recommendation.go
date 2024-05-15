package repositories

import (
	"context"

	"github.com/satryanararya/go-chefbot/entities"
	rec_const "github.com/satryanararya/go-chefbot/constants/recommendation"
	"gorm.io/gorm"
)

type RecommendationRepository interface {
	GetRecommendation(ctx context.Context, userID int64) (*[]entities.Recommendation, error)
	CreateRecommendation(ctx context.Context, recommendation *[]entities.Recommendation) error
}

type recommendationRepository struct {
	DB *gorm.DB
}

func NewRecommendationRepository(db *gorm.DB) *recommendationRepository {
	return &recommendationRepository{
		DB: db,
	}
}

func (r *recommendationRepository) GetRecommendation(ctx context.Context, userID int64) (*[]entities.Recommendation, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	
	recommendations := new([]entities.Recommendation)

	result := r.DB.Order("id DESC").Limit(rec_const.RECOMMENDATION_LIMIT).Where("user_id = ?", userID).Find(recommendations)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return recommendations, nil
}

func (r *recommendationRepository) CreateRecommendation(ctx context.Context, recommendation *[]entities.Recommendation) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	result := r.DB.WithContext(ctx).Create(recommendation)
	return result.Error
}