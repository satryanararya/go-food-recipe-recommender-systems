package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/repositories"
	"github.com/satryanararya/go-chefbot/utils/password"
	"github.com/satryanararya/go-chefbot/utils/token"
)

type UserUseCase interface {
	Register(c echo.Context, req *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
	Login(c echo.Context, req *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	GetUserByID(c echo.Context, id uuid.UUID) (*dto.UserGetByIDResponse, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
	passUtil password.PasswordUtil
	tokenUtil token.TokenUtil
}

func NewUserUseCase(ur repositories.UserRepository, pu password.PasswordUtil, tu token.TokenUtil) *userUseCase {
	return &userUseCase{
		userRepo: ur,
		passUtil: pu,
		tokenUtil: tu,
	}
}

func (uc *userUseCase) Register(c echo.Context, req *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	hashedPassword, err := uc.passUtil.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return &dto.UserRegisterResponse{
		Username: user.Username,
		Email: user.Email,
	}, uc.userRepo.CreateUser(ctx, user)
}

func (uc *userUseCase) Login(c echo.Context, req *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	user, err := uc.userRepo.GetUser(ctx, &entities.User{Email: req.Email})
	if err != nil {
		return nil, err
	}

	if err := uc.passUtil.ComparePassword(req.Password, user.Password); err != nil {
		return nil, err
	}

	token, err := uc.tokenUtil.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.UserLoginResponse{
		Username: user.Username,
		Email: user.Email,
		Token: token,
	}, nil
}

func (uc *userUseCase) GetUserByID(c echo.Context, id uuid.UUID) (*dto.UserGetByIDResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	user, err := uc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.UserGetByIDResponse{
		Username: user.Username,
		Email: user.Email,
		UserFoodPreference: &user.UserFoodPreference,
		UserCookingSkill: &user.UserCookingSkill,
		UserAllergies: &user.UserAllergies,
		Recipe: &user.Recipe,
		FavoriteRecipe: user.FavoriteRecipe,
	}, nil
}