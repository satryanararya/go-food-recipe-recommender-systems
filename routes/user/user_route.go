package user

import (
	"github.com/labstack/echo/v4"
	uc "github.com/satryanararya/go-chefbot/controllers/user"
	ur "github.com/satryanararya/go-chefbot/repositories"
	uuc "github.com/satryanararya/go-chefbot/usecases"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/satryanararya/go-chefbot/utils/password"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"github.com/satryanararya/go-chefbot/utils/token"
	"gorm.io/gorm"
)

func InitUserRoute(ug *echo.Group, db *gorm.DB, v *validation.Validator) {
	userRepository := ur.NewUserRepository(db)
	passUtil := password.NewPasswordUtil()
	tokenUtil := token.NewTokenUtil()

	userUseCase := uuc.NewUserUseCase(userRepository, passUtil, tokenUtil)
	userController := uc.NewUserController(userUseCase, v, tokenUtil)

	ug.POST("/register", userController.Register)
	ug.POST("/login", userController.Login)

	ug.Use(echojwt.WithConfig(token.GetJWTConfig()))
	ug.GET("/user", userController.GetUserByID)
}
