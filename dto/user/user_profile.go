package user

type UserFoodPreferencesRequest struct {
	DietaryRestriction *string `json:"dietary_restriction"`
	ReligiousReason    *string `json:"religious_reason"`
}

type UserCookingSkillRequest struct {
	ExperienceYears     string `json:"experience_years"`
	TimeCommitment      string `json:"time_commitment"`
	RecipeComplexity    string `json:"recipe_complexity"`
	IngredientDiversity string `json:"ingredient_diversity"`
}

//TODO: User Allergies
