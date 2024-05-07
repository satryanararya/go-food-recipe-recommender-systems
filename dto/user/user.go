package dto

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=25"`
	Email    string `json:"email" validate:"required,email"` 
	Password string `json:"password" validate:"required,min=8"`
}
