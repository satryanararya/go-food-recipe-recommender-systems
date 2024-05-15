package routes

import (
	"github.com/labstack/echo/v4"
	ratingreview "github.com/satryanararya/go-chefbot/routes/rating_review"
	favrecipe "github.com/satryanararya/go-chefbot/routes/favorite_recipe"
	"github.com/satryanararya/go-chefbot/routes/recipe"
	"github.com/satryanararya/go-chefbot/routes/user"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB, v *validation.Validator) {
	userRoute := e.Group("")
	userFoodPreferenceRoute := e.Group("/user-profile")
	userCookingSkillRoute := e.Group("/user-profile")
	userAllergiesRoute := e.Group("/user-profile")
	recipeRoute := e.Group("/recipe")
	ratingReviewRoute := e.Group("/recipe")
	favoriteRecipeRoute := e.Group("/recipe")
	recommendationRoute := e.Group("/recipe")

	user.InitUserRoute(userRoute, db, v)
	user.InitUserFoodPreferencesRoute(userFoodPreferenceRoute, db, v)
	user.InitUserCookingSkillRoute(userCookingSkillRoute, db, v)
	user.InitUserAllergiesRoute(userAllergiesRoute, db, v)
	recipe.InitRecipeRoute(recipeRoute, db, v)
	recipe.InitRecommendationRoute(recommendationRoute, db)
	ratingreview.InitRatingReviewRoute(ratingReviewRoute, db, v)
	favrecipe.InitFavoriteRecipeRoute(favoriteRecipeRoute, db, v)
}