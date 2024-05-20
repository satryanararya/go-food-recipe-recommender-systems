package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/satryanararya/go-chefbot/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetAllUsers(ctx context.Context) (*[]entities.User, error)
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

func (ur *userRepository) GetUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if err := ur.DB.Where(user).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	user := &entities.User{ID: id}
	err := ur.DB.Preload(clause.Associations).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetAllUsers(ctx context.Context) (*[]entities.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	users := new([]entities.User)
	if err := ur.DB.Preload(clause.Associations).Find(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}