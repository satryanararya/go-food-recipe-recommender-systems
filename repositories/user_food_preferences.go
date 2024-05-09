package repositories

import (
	"context"

	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type UserFoodPreferencesRepository interface {
	AddFoodPreference(ctx context.Context, userFoodPref *entities.UserFoodPreference) error
	EditFoodPreference(ctx context.Context, userFoodPref *entities.UserFoodPreference) error
    DeleteFoodPreference(ctx context.Context, userID int64) error
	GetFoodPreferenceByID(ctx context.Context, userID int64) (*entities.UserFoodPreference, error)
}

type userFoodPreferencesRepository struct {
	DB *gorm.DB
}

func NewUserFoodPreferencesRepository(db *gorm.DB) *userFoodPreferencesRepository {
	return &userFoodPreferencesRepository{
		DB: db,
	}
}

func(ur *userFoodPreferencesRepository) AddFoodPreference(ctx context.Context, userFoodPref *entities.UserFoodPreference) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return ur.DB.Create(userFoodPref).Error
}

func (ur *userFoodPreferencesRepository) EditFoodPreference(ctx context.Context, userFoodPref *entities.UserFoodPreference) error {
    if err := ctx.Err(); err != nil {
        return err
    }
    return ur.DB.Session(&gorm.Session{FullSaveAssociations: true}).Where("user_id = ?", userFoodPref.UserID).Updates(userFoodPref).Error
}

func (ur *userFoodPreferencesRepository) DeleteFoodPreference(ctx context.Context, userID int64) error {
    if err := ctx.Err(); err != nil {
        return err
    }
    return ur.DB.Where("user_id = ?", userID).Delete(&entities.UserFoodPreference{}).Error
}

// func (ur *userFoodPreferencesRepository) GetFoodPreferenceByID(ctx context.Context, userID int64) (*entities.UserFoodPreference, error) {
//     if err := ctx.Err(); err != nil {
//         return nil, err
//     }
    
//     pref := &entities.UserFoodPreference{}
//     if err := ur.DB.Where("user_id = ?", userID).First(pref).Error; err != nil {
//         return nil, err
//     }
    
//     return pref, nil
// }