package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"github.com/satryanararya/go-chefbot/routes/user"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB, v *validation.Validator) {
	userRoute := e.Group("")
	userFoodPreference := e.Group("/user-profile")

	user.InitUserRoute(userRoute, db, v)
	user.InitUserFoodPreferencesRoute(userFoodPreference, db, v)
}