package entity

type Post struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	UsersId    string `gorm:"type: varchar(300)" json:"users_id"`
	Post       string `gorm:"type:text" json:"post"`
	PictureUrl string `gorm:"type: varchar(200)" json:"picture_url"`
}
