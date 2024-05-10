package repositories

import (
    "context"

    "github.com/satryanararya/go-chefbot/entities"
    "gorm.io/gorm"
)

type UserAllergyRepository interface {
    Save(ctx context.Context, userID int64, ingredientID int64) error
    GetAllergies(ctx context.Context, userID int64) ([]*entities.UserAllergies, error)
}

type userAllergyRepository struct {
    DB *gorm.DB
}

func NewUserAllergyRepository(db *gorm.DB) UserAllergyRepository {
    return &userAllergyRepository{
        DB: db,
    }
}

func (r *userAllergyRepository) Save(ctx context.Context, userID int64, ingredientID int64) error {
    userAllergy := &entities.UserAllergies{
        UserID:       userID,
        IngredientID: ingredientID,
    }
    result := r.DB.WithContext(ctx).Save(userAllergy)
    return result.Error
}

func (r *userAllergyRepository) GetAllergies(ctx context.Context, userID int64) ([]*entities.UserAllergies, error) {
    var allergies []*entities.UserAllergies
    result := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&allergies)
    return allergies, result.Error
}