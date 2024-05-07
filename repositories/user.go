package repositories

import (
	"context"

	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *entities.User) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return ur.DB.Create(user).Error
}