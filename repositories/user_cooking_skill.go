package repositories

import (
	"context"

	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type UserCookingSkillRepository interface {
	AddCookingSkill(ctx context.Context, userCookingSkill *entities.UserCookingSkill) error
	EditCookingSkill(ctx context.Context, userCookingSkill *entities.UserCookingSkill) error
}

type userCookingSkillRepository struct {
	DB *gorm.DB
}

func NewUserCookingSkillRepository(db *gorm.DB) *userCookingSkillRepository {
	return &userCookingSkillRepository{
		DB: db,
	}
}

func (ur *userCookingSkillRepository) AddCookingSkill(ctx context.Context, userCookingSkill *entities.UserCookingSkill) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return ur.DB.Create(userCookingSkill).Error
}

func (ur *userCookingSkillRepository) EditCookingSkill(ctx context.Context, userCookingSkill *entities.UserCookingSkill) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return ur.DB.Session(&gorm.Session{FullSaveAssociations: true}).Where("user_id = ?", userCookingSkill.UserID).Updates(userCookingSkill).Error
}
