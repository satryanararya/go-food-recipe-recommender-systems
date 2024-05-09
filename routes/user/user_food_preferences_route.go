package user

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	uc "github.com/satryanararya/go-chefbot/controllers/user"
	ur "github.com/satryanararya/go-chefbot/repositories"
	uuc "github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

func InitUserFoodPreferencesRoute(ufpg *echo.Group, db *gorm.DB, v *validation.Validator) {
	userFoodPreferenceRepo := ur.NewUserFoodPreferencesRepository(db)
	userFoodPreferenceUseCase := uuc.NewUserFoodPreferencesUseCase(userFoodPreferenceRepo)
	tokenUtil := token.NewTokenUtil()

	userFoodPreferenceController := uc.NewUserFoodPreferencesController(userFoodPreferenceUseCase, v, tokenUtil)

	ufpg.Use(echojwt.WithConfig(token.GetJWTConfig()))
	ufpg.POST("/food-preferences", userFoodPreferenceController.AddUserFoodPreferences)
	ufpg.PUT("/food-preferences", userFoodPreferenceController.EditUserFoodPreferences)
	ufpg.DELETE("/food-preferences", userFoodPreferenceController.DeleteUserFoodPreferences)
}
