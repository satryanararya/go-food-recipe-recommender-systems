package usecases

import (
	"context"

	"github.com/labstack/echo/v4"
	dto "github.com/satryanararya/go-chefbot/dto/user"
	"github.com/satryanararya/go-chefbot/repositories"
	"github.com/satryanararya/go-chefbot/entities"
	"github.com/satryanararya/go-chefbot/utils/password"
)

type UserUseCase interface {
	Register(c echo.Context, req *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
	passUtil password.PasswordUtil
}

func NewUserUseCase(ur repositories.UserRepository, pu password.PasswordUtil) *userUseCase {
	return &userUseCase{
		userRepo: ur,
		passUtil: pu,
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
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return &dto.UserRegisterResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
	}, uc.userRepo.CreateUser(ctx, user)
}