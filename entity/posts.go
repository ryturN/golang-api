package entity

import "time"

type Post struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	UsersId    string `gorm:"type: varchar(300)" json:"users_id"`
	Post       string `gorm:"type:text" json:"post"`
	PictureUrl *string
	CreatedAt  time.Time `gorm:"type: timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type: timestamp" json:"updated_at"`
}
