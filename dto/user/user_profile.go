package user

type UserFoodPreferencesRequest struct {
	DietaryRestriction *string `json:"dietary_restriction"`
	ReligiousReason    *string `json:"religious_reason"`
}

type UserCookingSkillRequest struct {
	ExperienceYears     *int    `json:"experience_years"`
	TimeCommitment      *int    `json:"time_commitment"`
	RecipeComplexity    *string `json:"recipe_complexity"`
	IngredientDiversity *string `json:"ingredient_diversity"`
}

//TODO: User Allergies