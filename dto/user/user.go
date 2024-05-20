package user

import "github.com/satryanararya/go-chefbot/entities"

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserGetByIDResponse struct {
	Username           string                      `json:"username"`
	Email              string                      `json:"email"`
	UserFoodPreference *entities.UserFoodPreference `json:"food_preference"`
	UserCookingSkill   *entities.UserCookingSkill   `json:"cooking_skill"`
	UserAllergies      *[]entities.UserAllergies    `json:"allergies"`
	Recipe             *[]entities.Recipe           `json:"recipe"`
	FavoriteRecipe     *[]entities.FavoriteRecipe  `json:"favorite_recipe"`
}
