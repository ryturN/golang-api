package entity

type Users struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	UsersId  string `gorm:"type: varchar(300)" json:"users_id"`
	FullName string `gorm:"type:varchar(300)" json:"full_name"`
	Username string `gorm:"type:varchar(300)" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(300)" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(300)" json:"password" `
}
