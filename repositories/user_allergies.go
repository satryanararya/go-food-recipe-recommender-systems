package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type UserAllergyRepository interface {
	Save(ctx context.Context, userAllergy *entities.UserAllergies) error
	GetAllergies(ctx context.Context, userID uuid.UUID) ([]*entities.UserAllergies, error)
}

type userAllergyRepository struct {
	DB *gorm.DB
}

func NewUserAllergyRepository(db *gorm.DB) UserAllergyRepository {
	return &userAllergyRepository{
		DB: db,
	}
}

func (r *userAllergyRepository) Save(ctx context.Context, userAllergy *entities.UserAllergies) error {
	result := r.DB.WithContext(ctx).Create(userAllergy)
	return result.Error
}

func (r *userAllergyRepository) GetAllergies(ctx context.Context, userID uuid.UUID) ([]*entities.UserAllergies, error) {
	var allergies []*entities.UserAllergies
	result := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&allergies)
	return allergies, result.Error
}
