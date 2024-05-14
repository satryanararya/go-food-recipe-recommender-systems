package message

const (
	// user
	MsgUserCreated         = "user created successfully"
	MsgLoginSuccess        = "login successfully"
	MsgRetrieveUserSuccess = "user retreived successfully"
	MsgUpdateUserSuccess   = "user updated successfully"

	// user food preference
	MsgAddPreferenceSuccess = "food preference added successfully"
	MsgEditPreferenceSuccess = "food preference edited successfully"
	MsgDeletePreferenceSuccess = "food preference deleted successfully"
	MsgPreferenceUpdated         = "preference updated successfully"
	MsgRetrievePreferenceSuccess = "preference retreived successfully"

	// user cooking skill
	MsgAddCookingSkillSuccess = "cooking skill added successfully"
	MsgEditCookingSkillSuccess = "cooking skill edited successfully"

	// user allergy
	MsgAddUserAllergySuccess = "allergy added successfully"

	// recipe
	MsgGetRecipeSuccess = "recipe retreived successfully"
	MsgDeleteRecipeSuccess = "recipe deleted successfully"

	// rating and review
	MsgAddRatingReviewSuccess = "rating and review added successfully"
	MsgDeleteRatingReviewSuccess = "rating and review deleted successfully"
	MsgGetRatingReviewSuccess = "rating and review retreived successfully"

	// favorite recipe
	MsgAddFavoriteRecipeSuccess = "favorite recipe added successfully"
	MsgRemoveFavoriteRecipeSuccess = "favorite recipe removed successfully"
	MsgGetFavoriteRecipesSuccess = "favorite recipes retreived successfully"
)
