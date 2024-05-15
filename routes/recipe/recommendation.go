package recipe

import (
	"os"

	"github.com/labstack/echo/v4"
	recipeClient "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	openAIClient "github.com/satryanararya/go-chefbot/drivers/openai"
	"github.com/satryanararya/go-chefbot/repositories"
	"github.com/satryanararya/go-chefbot/usecases"
	rc "github.com/satryanararya/go-chefbot/controllers/recommendation"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/satryanararya/go-chefbot/utils/token"
	"gorm.io/gorm"
)

func InitRecommendationRoute(rcg *echo.Group, db *gorm.DB) {
	spoonacularAPIKey := os.Getenv("SPOONACULAR_API_KEY")
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")

	recommendationRepo := repositories.NewRecommendationRepository(db)
	userRepo := repositories.NewUserRepository(db)

	recipeClient := recipeClient.NewRecipeClient(spoonacularAPIKey)
	openAIClient := openAIClient.NewOpenAIClient(openAIAPIKey)

	tokenUtil := token.NewTokenUtil()

	recommendationUseCase := usecases.NewRecommendationUseCase(recommendationRepo, userRepo, recipeClient, openAIClient)

	recommendationUseCase.StartRecommendationCron()

	recommendationController := rc.NewRecommendationController(recommendationUseCase, tokenUtil)

	rcg.Use(echojwt.WithConfig(token.GetJWTConfig()))

	rcg.GET("/recommendations", recommendationController.GetRecommendation)
}