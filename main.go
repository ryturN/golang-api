package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/controllers/logincontroller"
	"github.com/golang-api/models"
	"github.com/golang-api/routers"
)

func main() {
	models.LoadConfig()
	models.Connected()

	router := gin.Default()
	v1 := router.Group("/v1")
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

	v3 := router.Group("/v3")
	v3.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Connected!",
		})
	})
	routers.AuthRouter(v3)
	router.Run(":3030")
}
