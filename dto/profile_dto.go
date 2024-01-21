package dto

type ProfileRequest struct {
	UserId string `json:"user_id"`
}
type ProfileResponse struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Url      any    `json:"url"`
}

type UpdateProfileRequest struct {
	UsersId  string `json:"users_id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fotos    string `json:"fotos"`
}
