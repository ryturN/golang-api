package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connected establishes a connection to the database
func Connected() {
	database, err := gorm.Open(mysql.Open("root:ryan14@tcp(localhost:3306)/golang_db"))
	if err != nil {
		panic(err.Error())
	}
	database.AutoMigrate(&Users{})
	DB = database

}
