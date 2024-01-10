package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/controllers/logincontroller"
	"github.com/golang-api/models"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	models.Connected()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	v1.POST("/login", logincontroller.Login)
	v1.GET("/test", logincontroller.Test)
	v1.POST("/register", logincontroller.Register)
	v1.GET("/findUser", logincontroller.FindUser)
	v1.GET("/findUser/:id", logincontroller.FindUserById)

	v2 := router.Group("/v2")
	v2.POST("/login", logincontroller.Login)
	v2.GET("/test", logincontroller.Test)
	v2.POST("/register", logincontroller.Register)
	v2.GET("/findUser", logincontroller.FindUser)
	v2.GET("/findUser/:id", logincontroller.FindUserById)

	router.Run("localhost:3030")
}
