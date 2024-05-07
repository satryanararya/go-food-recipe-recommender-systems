package user

import (
	"github.com/labstack/echo/v4"
	uc "github.com/satryanararya/go-chefbot/controllers/user"
	ur "github.com/satryanararya/go-chefbot/repositories"
	uuc "github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/password"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

func InitUserRoute(ug *echo.Group, db *gorm.DB, v *validation.Validator) {
	userRepository := ur.NewUserRepository(db)
	passUtil := password.NewPasswordUtil()

	userUseCase := uuc.NewUserUseCase(userRepository, passUtil)
	userController := uc.NewUserController(userUseCase, v)

	ug.POST("/register", userController.Register)
}
