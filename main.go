package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logincontroller "github.com/golang-api/controllers/productcontroller"
	"github.com/golang-api/models"
)

func main() {
	router := gin.Default()
	models.Connected()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	router.POST("/login", logincontroller.Login)
	router.GET("/test", logincontroller.Test)
	router.POST("/register", logincontroller.Register)
	router.GET("/findUser", logincontroller.FindUser)
	router.GET("/findUser/:id", logincontroller.FindUserById)
	router.Run("localhost:3030")
}
