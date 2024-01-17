package dto

import "mime/multipart"

type PostResponse struct {
	Id         int    `json:"id"`
	UsersId    string `json:"users_id"`
	User       User   `gorm:"foreignKey:UsersId" json:"users_id"`
	Post       string `gorm:"text" json:"post"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type PostRequest struct {
	UsersId string                `form:"users_id"`
	Post    string                `form:"post"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
