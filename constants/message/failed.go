package message

const (
	// file
	MsgFailedOpenFile = "failed to open file"

	// password
	MsgFailedHashingPassword = "failed hashing password"
	MsgPasswordMismatch      = "password mismatch"

	// request
	MsgMismatchedDataType = "mismatched data type"
	MsgInvalidRequestData = "invalid request data"

	// database
	MsgFailedConnectDB = "failed connect to database"
	MsgSeedFailed      = "database seeding failed"
	MsgFailedMigrateDB = "failed to migrate database"

	// users
	MsgUserCreationFailed = "failed to create user"
	MsgUserExist          = "username or email already exist"
	MsgLoginFailed        = "failed to login"
	MsgUnregisteredEmail  = "unregistered email"
	MsgUnregisteredUser   = "unregistered user"
	MsgGetUserFailed      = "failed to retreive user"
	MsgUpdateUserFailed   = "failed to update user"

	// user food preference
	MsgAddPreferenceFailed = "failed to add your food preference"
	MsgEditPreferenceFailed  = "failed to update your food preference"
	MsgDeletePreferenceFailed = "failed to delete your food preference"
	MsgPreferenceInputNotFound = "preference types not found"

	// user cooking skill
	MsgAddCookingSkillFailed = "failed to add your cooking skill"
	MsgEditCookingSkillFailed  = "failed to update your cooking skill"

	// user allergies
	MsgAddUserAllergyFailed = "failed to add your allergy"

	// recipe
	MsgInvalidID = "invalid recipe ID"
	MsgGetRecipeFailed = "failed to get recipe"
	MsgDeleteRecipeFailed = "failed to delete recipe"
	MsgUpdateRecipeFailed = "failed to update recipe"

	// rating and review
	MsgAddRatingReviewFailed = "failed to add rating and review"
	MsgDeleteRatingReviewFailed = "failed to delete rating and review"
	MsgGetRatingReviewFailed = "failed to retrieve rating and review"

	// favorite recipe
	MsgAddFavoriteRecipeFailed = "failed to add favorite recipe"
	MsgRemoveFavoriteRecipeFailed = "failed to remove favorite recipe"
	MsgGetFavoriteRecipesFailed = "failed to get favorite recipes"

	// recommendation
	MsgGetRecommendationFailed = "failed to get recommendation"
	MsgFailedAddRecommendationCron = "failed to add recommendation cron"
	MsgFailedCreateRecommendation  = "failed to create recommendation"
	MsgRecommendationNotFound      = "recommendation not found"

	// token
	MsgFailedGeneratingToken = "failed generating token"
	MsgUnauthorized          = "unauthorized user"
	MsgInvalidToken          = "invalid token"

	MsgExternalServiceError   = "external service error"
	MsgQueryMinimum           = "item must at least 3 characters"
	MsgSearchItemFailed       = "failed to search item"
	MsgItemNotFound           = "item not found"
	MsgGetItemNutritionFailed = "failed to get item nutrition"

	// pages
	MsgPageNotFound = "page not found"
)
