package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron"
	msg "github.com/satryanararya/go-chefbot/constants/message"
	openAIClient "github.com/satryanararya/go-chefbot/drivers/openai"
	recipeClient "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	"github.com/satryanararya/go-chefbot/repositories"
	"github.com/satryanararya/go-chefbot/utils/prompt"
	rec_util "github.com/satryanararya/go-chefbot/utils/recommendation"
)

type RecommendationUseCase interface {
	GetRecommendation(c echo.Context, userID int64) (*[]recipeClient.RecipeInformation, error)
}

type recommendationUseCase struct {
	recommendationRepo repositories.RecommendationRepository
	userRepo           repositories.UserRepository

	recipeClient recipeClient.RecipeClient
	openAIClient openAIClient.OpenAIClient
}

func NewRecommendationUseCase(rr repositories.RecommendationRepository, ur repositories.UserRepository, rc recipeClient.RecipeClient, oai openAIClient.OpenAIClient) *recommendationUseCase {
	return &recommendationUseCase{
		recommendationRepo: rr,
		userRepo:           ur,
		recipeClient:       rc,
		openAIClient:       oai,
	}
}

func (r *recommendationUseCase) StartRecommendationCron() {
	c := cron.New()
	if err := c.AddFunc("@weekly", r.FetchOpenAIRecommendation); err != nil {
		log.Fatal(msg.MsgFailedAddRecommendationCron)
	}

	go func() {
		c.Start()
		defer c.Stop()

		select {}
	}()
}

func (r *recommendationUseCase) GetRecommendation(c echo.Context, userID int64) (*[]recipeClient.RecipeInformation, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	recommendations, err := r.recommendationRepo.GetRecommendation(ctx, userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var recipeInformations []recipeClient.RecipeInformation
	for _, recommendation := range *recommendations {
		searchRecipes, err := r.recipeClient.SearchRecipe(ctx, recommendation.RecipeName)
		fmt.Println("searchRecipes: ", searchRecipes)
		if err != nil {
			return nil, err
		}
		for _, result := range searchRecipes.Results {
			recipeInformation, err := r.recipeClient.GetRecipeInformation(ctx, int(result.ID))
			if err != nil {
				return nil, err
			}
			recipeInformations = append(recipeInformations, recipeInformation)
		}
	}
	return &recipeInformations, nil
}

func (r *recommendationUseCase) FetchOpenAIRecommendation() {
	ctx := context.Background()

	users, _ := r.userRepo.GetAllUsers(ctx)
	for _, user := range *users {
		prompt := prompt.GetRecommendationPrompt(&user)
		result, err := r.openAIClient.GetRecommendation(prompt, rec_util.ToString(user.Recommendations, true))
		fmt.Println("result:", result)
		if err == nil {
			recommendations := rec_util.ToStruct(result, user.ID)
			fmt.Println("recommendations: ", recommendations)
			if err := r.recommendationRepo.CreateRecommendation(ctx, recommendations); err != nil {
				log.Println(msg.MsgFailedCreateRecommendation)
			}
		}

	}
}

// users, err := r.userRepo.GetAllUser(ctx)
// if err != nil {
// 	log.Println(msg.MsgFailedGetAllUser)
// 	return
// }

// for _, user := range *users {
// 	openAIResponse, err := r.openAIClient.GetOpenAIResponse(ctx, user.ID)
// 	if err != nil {
// 		log.Println(msg.MsgFailedGetOpenAIResponse)
// 		continue
// 	}

// 	recommendation := rec_util.ParseOpenAIResponse(openAIResponse)

// 	if err := r.recommendationRepo.CreateRecommendation(ctx, recommendation); err != nil {
// 		log.Println(msg.MsgFailedCreateRecommendation)
// 		continue
// 	}
// }
