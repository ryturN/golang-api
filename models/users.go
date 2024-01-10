package models

type Users struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(300)" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(300)" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(300)" json:"password" `
}
