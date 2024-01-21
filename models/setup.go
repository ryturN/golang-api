package models

import (
	"fmt"
	"log"

	"github.com/golang-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connected establishes a connection to the database
func Connected() {

	connectStr := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_DATABASE)
	fmt.Println(connectStr)
	database, err := gorm.Open(mysql.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic("fail to connect")
	}
	log.Printf("Connected to database")
	database.AutoMigrate(entity.Users{})
	database.AutoMigrate(&entity.Foto{})
	database.AutoMigrate(&entity.Post{})
	DB = database

}
