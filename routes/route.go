package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"github.com/satryanararya/go-chefbot/routes/user"
	"github.com/satryanararya/go-chefbot/routes/recipe"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB, v *validation.Validator) {
	userRoute := e.Group("")
	userFoodPreferenceRoute := e.Group("/user-profile")
	userCookingSkillRoute := e.Group("/user-profile")
	userAllergiesRoute := e.Group("/user-profile")
	recipeRoute := e.Group("/recipe")

	user.InitUserRoute(userRoute, db, v)
	user.InitUserFoodPreferencesRoute(userFoodPreferenceRoute, db, v)
	user.InitUserCookingSkillRoute(userCookingSkillRoute, db, v)
	user.InitUserAllergiesRoute(userAllergiesRoute, db, v)
	recipe.InitRecipeRoute(recipeRoute, db, v)
}