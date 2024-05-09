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

func InitUserCookingSkillRoute(ufpg *echo.Group, db *gorm.DB, v *validation.Validator) {
	userCookingSkillRepo := ur.NewUserCookingSkillRepository(db)
	userCookingSkillUseCase := uuc.NewUserCookingSkillUseCase(userCookingSkillRepo)
	tokenUtil := token.NewTokenUtil()

	userCookingSkillController := uc.NewUserCookingSkillController(userCookingSkillUseCase, v, tokenUtil)

	ufpg.Use(echojwt.WithConfig(token.GetJWTConfig()))
	ufpg.POST("/cooking-skill", userCookingSkillController.AddUserCookingSkill)
	ufpg.PUT("/cooking-skill", userCookingSkillController.EditUserCookingSkill)
}