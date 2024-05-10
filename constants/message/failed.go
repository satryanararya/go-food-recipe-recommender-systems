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

	// token
	MsgFailedGeneratingToken = "failed generating token"
	MsgUnauthorized          = "unauthorized user"
	MsgInvalidToken          = "invalid token"

	MsgExternalServiceError   = "external service error"
	MsgQueryMinimum           = "item must at least 3 characters"
	MsgSearchItemFailed       = "failed to search item"
	MsgItemNotFound           = "item not found"
	MsgGetItemNutritionFailed = "failed to get item nutrition"
)
