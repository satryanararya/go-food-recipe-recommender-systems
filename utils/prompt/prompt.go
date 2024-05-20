package prompt

import (
	"fmt"

	"github.com/satryanararya/go-chefbot/entities"
)

func GetRecommendationPrompt(u *entities.User) string {
	start := "You are recipe recommender and give five list of best common recipe based on this preference:\n"

	pref := "- Dietary restriction : %s\n- Religious reason : %s\n- Experience years : %s\n- Time commitment : %s\n- Recipe complexity : %s\n- Ingredient diversity : %s\n- Allergies : %s\n- Favorite recipe : %s\n\n"

	var (
		dietaryRestriction  = "any"
		religiousReason     = "halal"
		experienceYears     = "1"
		timeCommitment      = "1"
		recipeComplexity    = "easy"
		ingredientDiversity = "rarely"
		allergies           = "none"
		favoriteRecipe      = "-"
	)

	if u.UserFoodPreference.DietaryRestriction != nil {
		dietaryRestriction = *u.UserFoodPreference.DietaryRestriction
	}
	if u.UserFoodPreference.ReligiousReason != nil {
		religiousReason = *u.UserFoodPreference.ReligiousReason
	}
	if u.UserCookingSkill.ExperienceYears != "" {
		experienceYears = u.UserCookingSkill.ExperienceYears
	}
	if u.UserCookingSkill.TimeCommitment != "" {
		timeCommitment = u.UserCookingSkill.TimeCommitment
	}
	if u.UserCookingSkill.RecipeComplexity != "" {
		recipeComplexity = u.UserCookingSkill.RecipeComplexity
	}
	if u.UserCookingSkill.IngredientDiversity != "" {
		ingredientDiversity = u.UserCookingSkill.IngredientDiversity
	}
	if len(u.UserAllergies) != 0 {
		allergies = u.UserAllergies[0].Ingredient.Name
	}
	if len(*u.FavoriteRecipe) != 0 {
		favoriteRecipe = (*u.FavoriteRecipe)[0].RecipeTitle
	}
	pref = fmt.Sprintf(pref, dietaryRestriction, religiousReason, experienceYears, timeCommitment, recipeComplexity, ingredientDiversity, allergies, favoriteRecipe)

	end := "\n*Response specification : Your recipe recommendation name should consist only 1 word (ex 'spaghetti', 'zucchini', 'turkey'). only name without any words at the beginning.*"

	prompt := fmt.Sprint(start, pref, end)

	return prompt
}
