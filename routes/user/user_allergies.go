package user

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	uc "github.com/satryanararya/go-chefbot/controllers/user"
	ur "github.com/satryanararya/go-chefbot/repositories"
	uuc "github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"github.com/satryanararya/go-chefbot/drivers/spoonacular_api/ingredients"
	"gorm.io/gorm"
)

func InitUserAllergiesRoute(uag *echo.Group, db *gorm.DB, v *validation.Validator) {
	apiKey := os.Getenv("SPOONACULAR_API_KEY")

	ingredientClient := ingredients.NewIngredientClient(apiKey)
	userAllergiesRepo := ur.NewUserAllergyRepository(db)
	ingredientsRepo := ur.NewIngredientRepository(db)

	userAllergiesUseCase := uuc.NewUserAllergiesUseCase(ingredientClient, ingredientsRepo, userAllergiesRepo)
	tokenUtil := token.NewTokenUtil()

	userAllergiesController := uc.NewUserAllergiesController(userAllergiesUseCase, v, tokenUtil)

	uag.Use(echojwt.WithConfig(token.GetJWTConfig()))
	uag.POST("/allergies", userAllergiesController.AddUserAllergies)
}
