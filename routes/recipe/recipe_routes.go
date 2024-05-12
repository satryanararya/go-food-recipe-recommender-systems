package recipe

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	rc "github.com/satryanararya/go-chefbot/controllers/recipe"
	client "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	"github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
)

func InitRecipeRoute(rg *echo.Group, v *validation.Validator) {
	apiKey := os.Getenv("SPOONACULAR_API_KEY")

	recipeClient := client.NewRecipeClient(apiKey)
	recipeUseCase := usecases.NewRecipeUseCase(recipeClient)
	recipeController := rc.NewRecipeController(recipeUseCase, v)

	rg.Use(echojwt.WithConfig(token.GetJWTConfig()))

	rg.GET("/search", recipeController.SearchRecipe)
	rg.GET("/info/:id", recipeController.GetRecipeInformation)
}