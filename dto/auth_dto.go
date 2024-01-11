package dto

type RegisterRequest struct {
	FullName        string `json:"full_name"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"password_confirmation"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
