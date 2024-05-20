package recipe

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	rc "github.com/satryanararya/go-chefbot/controllers/recipe"
	cs "github.com/satryanararya/go-chefbot/drivers/cloudinary"
	client "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	rr "github.com/satryanararya/go-chefbot/repositories"
	ru "github.com/satryanararya/go-chefbot/usecases"
	"github.com/satryanararya/go-chefbot/config"
	"github.com/satryanararya/go-chefbot/utils/token"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

func InitRecipeRoute(rg *echo.Group, db *gorm.DB, v *validation.Validator) {
    apiKey := os.Getenv("SPOONACULAR_API_KEY")

    recipeClient := client.NewRecipeClient(apiKey)
    cloudinaryInstance, _ := config.SetupCloudinary()
    cloudinaryService := cs.NewCloudinaryService(cloudinaryInstance)
    recipeRepo := rr.NewRecipeRepository(db)
    recipeUseCase := ru.NewRecipeUseCase(recipeClient, recipeRepo, cloudinaryService)
    
	tokenUtil := token.NewTokenUtil()
	
	recipeController := rc.NewRecipeController(recipeUseCase, v, tokenUtil)

    rg.Use(echojwt.WithConfig(token.GetJWTConfig()))

    rg.GET("", recipeController.GetUserRecipes)
    rg.GET("/search", recipeController.SearchRecipe)
    rg.GET("/info/:id", recipeController.GetRecipeInformation)
    rg.POST("/create", recipeController.CreateRecipe)
	rg.POST("/:recipeID/image", recipeController.UploadRecipeImage)
	rg.PUT("/:recipeID", recipeController.UpdateRecipe)
	rg.DELETE("/:recipeID", recipeController.DeleteRecipe)
}