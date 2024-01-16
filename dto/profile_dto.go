package dto

type ProfileRequest struct {
	UserId string `json:"user_id"`
}
type ProfileResponse struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Url      string `json:"url"`
}
