package dto

type AuthBodyDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
