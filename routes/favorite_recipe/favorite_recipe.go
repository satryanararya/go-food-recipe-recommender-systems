package favoriterecipe

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	fc "github.com/satryanararya/go-chefbot/controllers/favorite_recipe"
	fr "github.com/satryanararya/go-chefbot/repositories"
	fu "github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

// InitFavoriteRecipeRoute initializes favorite recipe routes
func InitFavoriteRecipeRoute(frg *echo.Group, db *gorm.DB, v *validation.Validator) {
	favoriteRecipeRepo := fr.NewFavoriteRecipeRepository(db)
	favoriteRecipeUseCase := fu.NewFavoriteRecipeUsecase(favoriteRecipeRepo)

	tokenUtil := token.NewTokenUtil()

	favoriteRecipeController := fc.NewFavoriteRecipeController(favoriteRecipeUseCase, v, tokenUtil)

	frg.Use(echojwt.WithConfig(token.GetJWTConfig()))

	frg.POST("/favorite_recipe/:id", favoriteRecipeController.AddToFavorites)
	frg.DELETE("/favorite_recipe/:id", favoriteRecipeController.RemoveFromFavorites)
	frg.GET("/favorites", favoriteRecipeController.FindFavoritesByUserID)
}